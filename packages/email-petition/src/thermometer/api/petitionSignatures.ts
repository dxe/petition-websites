//add petition thermometer implementation here here

const API_URL = "https://helptheducks.dxe.io/tally?campaign=";

export const fetchPetitionSignatureCountData = async ({
  startDate,
  campaignName,
}: {
  startDate: string;
  campaignName: string;
}) => {
  const resp = await fetch(`${API_URL}${campaignName}?start_date=${startDate}`);
  const json = (await resp.json()) as { total?: number | null };
  return {
    total: json.total ?? 0,
  };
};
