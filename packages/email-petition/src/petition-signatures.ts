const API_URL_BASE = "https://helptheducks.dxe.io/tally?campaign=";

export const fetchPetitionSignatureCount = async ({
  campaignName,
}: {
  campaignName: string;
}) => {
  const resp = await fetch(`${API_URL_BASE}${campaignName}`);
  const json = (await resp.json()) as { total?: number | null };
  return {
    total: json.total ?? 0,
  };
};
