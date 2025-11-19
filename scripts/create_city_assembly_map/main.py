#!/usr/bin/env python3
#
# Script takes in a number of municipality data files and outputs a mapping from municipality
# in CA to state assembly district, by percentage of the municipality's population that lives
# there. This script was written by ChatGPT and the sources found were on the instruction of
# ChatGPT, but the output looks accurate
#
# Sources fetched most recently Nov 2025, and script run in Nov 2025
#
# Usage:
#   python main.py <optional places csv name> <optional json map object file name>
#
# Creates the following file outputs:
#   ca_assembly_places.csv
#       CSV with columns:
#           district: State Assembly district number
#           municipality: Name of city or CDP
#           population: Population of the municipality in that district
#           percent_in_place: Percentage of the municipality's population in that district
#           district_share_percent: Percentage of the district's population in that municipality
#           place_type: "C" for incorporated city, "U" for CDP
#   ca_assembly_places_city_mapping.txt
#       Text file with a JSON-like object mapping each city/CDP name to a dictionary of
#           district: State Assembly district number
#           population: Population of the municipality in that district
#           percent_in_place: Percentage of the municipality's population in that district
#           district_share_percent: Percentage of the district's population in that municipality
#
# Expects these files in the working directory:
#   06_CA_SLDL22.txt
#       Fetched Nov 2025
#       Contains: 2020 Census blocks → 2022 State Assembly districts
#       Link: https://www2.census.gov/programs-surveys/decennial/rdo/mapping-files/2023/2022-state-legislative-bef/sldl_2022.zip
#   BlockAssign_ST06_CA/
#       Fetched Nov 2025
#       Contains: 2020 Census blocks → INCPLACEFP/CDPFP
#       Link: https://www2.census.gov/geo/docs/maps-data/data/baf/
#   NAMES_ST06_CA/
#       Fetched Nov 2025
#       Contains: INCPLACEFP/CDPFP → place name
#       Link: https://www2.census.gov/geo/docs/maps-data/data/nlt/NLT_2020_ST06_CA.zip
#   blocks_ca_pl_cache.csv
#       Fetched Nov 2025
#       Contains: Mapping 2020 Census blocks to population
#       Probably found: https://www.census.gov/programs-surveys/decennial-census/about/rdo/summary-files.html

import os, re, sys
import pandas as pd
import numpy as np

STATE = "06"  # California

# Hardcoded input data files
BAF_DIR = "BlockAssign_ST06_CA"
NLT_DIR = "NAMES_ST06_CA"

SLDL_PATH = "06_CA_SLDL22.txt"
PLACE_BAF_PATH = os.path.join(BAF_DIR, "BlockAssign_ST06_CA_INCPLACE_CDP.txt")
NLT_INC_PATH = os.path.join(NLT_DIR, "NAMES_ST06_CA_INCPLACE.txt")
NLT_CDP_PATH = os.path.join(NLT_DIR, "NAMES_ST06_CA_CDP.txt")
BLOCKS_CACHE_PATH = "blocks_ca_pl_cache.csv"

# Default output file names
OUT_CSV_DEFAULT = "ca_assembly_places.csv"
OUT_TXT_DEFAULT = "ca_assembly_places_city_mapping.txt"

# ---------- Local readers ----------
def read_local_baf_sldl(path=SLDL_PATH):
    if not os.path.exists(path):
        raise FileNotFoundError(f"Missing SLDL BAF: {path}")
    # Try comma-separated (BEF format) first, then pipe-separated (BAF format)
    try:
        sldl = pd.read_csv(path, sep=",", dtype=str, engine="python")
    except:
        sldl = pd.read_csv(path, sep="|", dtype=str, engine="python")
    # Block ID column: GEOID or BLOCKID
    block_col = next((c for c in ("GEOID","BLOCKID") if c in sldl.columns), None)
    if not block_col:
        raise FileNotFoundError("SLDL file missing GEOID or BLOCKID column")
    # District column: SLDLST, DISTRICT, SLDL, SLDL20, SLDL16
    dist_col = next((c for c in ("SLDLST","DISTRICT","SLDL","SLDL20","SLDL16") if c in sldl.columns), None)
    if not dist_col:
        raise FileNotFoundError("SLDL file missing district column (SLDLST/DISTRICT/SLDL/SLDL20/SLDL16)")
    sldl = sldl[[block_col, dist_col]].rename(columns={block_col:"GEOID20", dist_col:"SLDL"})
    return sldl

def read_local_baf_place(path=PLACE_BAF_PATH):
    if not os.path.exists(path):
        raise FileNotFoundError(f"Missing INCPLACE/CDP BAF: {path}")
    plc = pd.read_csv(path, sep="|", dtype=str, engine="python")
    if "BLOCKID" not in plc.columns:
        raise FileNotFoundError("INCPLACE_CDP BAF missing BLOCKID column")
    key = next((c for c in ("PLACEFP","INCPLACEFP","CDPFP","PLACE") if c in plc.columns), None)
    if not key:
        raise FileNotFoundError("INCPLACE_CDP BAF missing PLACEFP/INCPLACEFP/CDPFP/PLACE column")
    plc = plc[["BLOCKID", key]].rename(columns={"BLOCKID":"GEOID20", key:"PLACEFP"})
    plc["PLACEFP"] = plc["PLACEFP"].astype(str).str.zfill(5)
    return plc

def read_local_nlt_places(inc_path=NLT_INC_PATH, cdp_path=NLT_CDP_PATH):
    for p in (inc_path, cdp_path):
        if not os.path.exists(p):
            raise FileNotFoundError(f"Missing NLT file: {p}")
    def norm(path, kind):
        df = pd.read_csv(path, sep="|", dtype=str, engine="python")
        key = "PLACEFP" if "PLACEFP" in df.columns else ("PLACE" if "PLACE" in df.columns else None)
        name_col = "NAME" if "NAME" in df.columns else ("NAMELSAD" if "NAMELSAD" in df.columns else None)
        if key is None or name_col is None:
            raise FileNotFoundError(f"{os.path.basename(path)} missing PLACEFP/PLACE or NAME/NAMELSAD")
        out = df[[key, name_col]].copy()
        out.rename(columns={key:"PLACEFP", name_col:"placename"}, inplace=True)
        out["place_type"] = "C" if kind=="inc" else "U"
        out["PLACEFP"] = out["PLACEFP"].astype(str).str.zfill(5)
        return out
    inc = norm(inc_path, "inc")
    cdp = norm(cdp_path, "cdp")
    plc = pd.concat([inc, cdp], ignore_index=True).drop_duplicates(subset="PLACEFP")
    return plc

def _parse_dist(x):
    m = re.search(r"(\d+)", str(x))
    return int(m.group(1)) if m else np.nan

def write_city_to_district_mapping(out, out_txt):
    """Write a text file mapping each city to its district(s) with population percentages."""
    # Filter to only cities and CDPs (exclude county remainders)
    # Also filter out rows with NaN percent values
    cities = out[out["place_type"].isin(["C", "U"]) & pd.notna(out["percent"])].copy()
    
    # Group by municipality
    city_groups = cities.groupby("municipality")
    
    with open(out_txt, "w") as f:
        f.write("{\n")
        
        for idx, (city, group) in enumerate(city_groups):
            # Sort by district
            group = group.sort_values("district")
            
            # Write city name
            f.write(f'\t"{city}": {{\n')
            
            # Write district -> percent mappings
            for _, row in group.iterrows():
                district = int(row["district"])
                percent = int(round(row["percent"]))
                f.write(f'\t\t{district}: {percent},\n')
            
            f.write('\t},\n')
            f.write('\n')
        
        f.write("}\n")

# ---------- Main ----------
def main(out_csv=OUT_CSV_DEFAULT, out_txt=None):
    # Read local mapping tables
    sldl = read_local_baf_sldl()
    plcmap = read_local_baf_place()
    nlt_places = read_local_nlt_places()

    # Blocks: use cached file
    if os.path.exists(BLOCKS_CACHE_PATH):
        blocks = pd.read_csv(BLOCKS_CACHE_PATH, dtype={"GEOID20":str,"pop":int})
    else:
        raise FileNotFoundError(f"Missing cached blocks file: {BLOCKS_CACHE_PATH}")

    # Join: blocks → district + place
    df = blocks.merge(sldl, on="GEOID20", how="left")
    if df["SLDL"].isna().any():
        missing = int(df["SLDL"].isna().sum())
        raise AssertionError(f"{missing} blocks missing SLDL mapping; check BAF file")
    df = df.merge(plcmap, on="GEOID20", how="left")
    df = df.merge(nlt_places, on="PLACEFP", how="left")

    # Filter out blocks without a place (no remainders)
    df = df[df["placename"].notna()]

    # District number as int
    df["district"] = df["SLDL"].apply(_parse_dist).astype(int)

    # ---------- Aggregation ----------
    # By district × place (keep PLACEFP to compute % of place)
    by = ["district","placename","place_type","PLACEFP"]
    agg = df.groupby(by, dropna=False)["pop"].sum().reset_index()

    # % of district total (district_share_percent)
    ad_tot = agg.groupby("district")["pop"].sum().rename("ad_total")
    agg = agg.merge(ad_tot, on="district", how="left")
    agg["district_share_percent"] = (agg["pop"] / agg["ad_total"] * 100).round(1)

    # Denominators for % of place
    place_tot = df.groupby("PLACEFP")["pop"].sum().rename("place_total")
    agg = agg.merge(place_tot, on="PLACEFP", how="left")

    # % of the place that lies in the district
    agg["percent_in_place"] = np.where(
        pd.notna(agg["place_total"]) & (agg["place_total"] > 0),
        (agg["pop"] / agg["place_total"] * 100).round(1),
        np.nan
    )

    # ---------- Output ----------
    out = agg.rename(columns={
        "placename":"municipality",
        "pop":"population"
    })[[
        "district",
        "municipality",
        "population",
        "percent_in_place",
        "district_share_percent",
        "place_type"
    ]]

    # Rename for convenience so 'percent' mirrors scraped pages
    out = out.rename(columns={"percent_in_place":"percent"})

    out.sort_values(["district","municipality"], inplace=True)
    out.to_csv(out_csv, index=False)
    print(f"Wrote {out_csv} with {len(out):,} rows.")

    # Write city-to-district mapping text file
    if out_txt is None:
        out_txt = out_csv.replace(".csv", "_city_mapping.txt")
    write_city_to_district_mapping(out, out_txt)
    print(f"Wrote {out_txt} with city-to-district mappings.")

    # QA summaries
    # Sum of district_share_percent per district should be ~100.0
    chk_d = out.groupby("district")["district_share_percent"].sum().round(1)
    print("Sum of district_share_percent per district (≈100.0 expected):")
    print(chk_d.describe())

    # For fully contained places, 'percent' should be ~100.0
    fully = out[(out["place_type"].isin(["C","U"])) & (out["percent"] >= 99.9)]
    print(f"Fully-contained place rows at ~100%: {len(fully):,}")

if __name__ == "__main__":
    out_csv = sys.argv[1] if len(sys.argv) > 1 else OUT_CSV_DEFAULT
    out_txt = sys.argv[2] if len(sys.argv) > 2 else None
    main(out_csv, out_txt)
