import { useQuery } from "@tanstack/react-query";
import React, { useMemo } from "react";
import classNames from "classnames";
import { fetchPetitionSignatureCount } from "./api";
import { getNextGoal } from "./calculate-goal";

const queryOptions = {
  // Refetch every minute.
  refetchInterval: 1000 * 60,
};

export const Thermometer = ({
  goal,
  campaignName,
}: {
  goal: number;
  campaignName: string;
}) => {
  const { data, isLoading, isError } = useQuery<{ total: number }>({
    queryKey: ["thermometer", campaignName],
    queryFn: () => fetchPetitionSignatureCount({ campaignName }),
    ...queryOptions,
  });
  const calculatedAmt = useMemo(() => data?.total ?? 0, [data?.total]);
  const calculatedGoal = useMemo(
    () => getNextGoal(calculatedAmt, goal),
    [calculatedAmt, goal],
  );
  const progress = useMemo(
    () => (!calculatedAmt ? 0 : (calculatedAmt / calculatedGoal) * 100),
    [calculatedAmt, calculatedGoal],
  );

  return isError ? null : (
    <div
      className={classNames(
        "w-full bg-white rounded-lg flex flex-col py-6 px-4 shadow-md gap-4",
        { "animate-pulse": isLoading },
      )}
    >
      <div className="flex justify-between gap-4">
        <div
          className={classNames(
            "flex flex-col md:flex-row gap-1 md:gap-3 md:items-center",
            { "text-gray-400": !data },
          )}
        >
          <span className="text-2xl font-medium">
            {data ? calculatedAmt : "✍️"}
          </span>
          <span className="text-sm whitespace-nowrap">
            {data?.total.toLocaleString() ?? "Loading"} emails sent
          </span>
        </div>
        <div className="self-end text-right">
          <span className="text-sm uppercase">Goal</span>{" "}
          <span className="font-medium">{calculatedGoal}</span>
        </div>
      </div>
      <div className="w-full bg-gray-200 rounded-full h-5 overflow-hidden">
        <div
          className="bg-blue-500 h-5 transition-all"
          style={{ width: `${progress}%` }}
        ></div>
      </div>
    </div>
  );
};
