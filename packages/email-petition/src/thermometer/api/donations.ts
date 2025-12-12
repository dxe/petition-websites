//add petition thermometer implementation here here

const API_URL = "https://helptheducks.dxe.io/tally?campaign=freezoe";

export const fetchDonationData = async ({
  startDate,
  campaignName,
}: {
  startDate: string;
  campaignName: string;
}) => {
  const resp = await fetch(`${API_URL}`); //await fetch(`${API_URL}${campaignName}?start_date=${startDate}`);
  const json = (await resp.json()) as { amt: string; count: string };
  return {
    amt: parseFloat(json.amt === "None" ? "0" : json.amt),
    count: parseFloat(json.count),
  };
};
