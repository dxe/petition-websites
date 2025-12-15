const INCREMENT = 10000;
const THRESHOLD = 1000;

export const getNextGoal = (amt: number, default_goal: number) => {
  if (amt < default_goal - THRESHOLD) {
    return default_goal;
  }

  const remainder = amt % INCREMENT;
  const nextGoal = amt - remainder + INCREMENT;

  return nextGoal - amt <= THRESHOLD ? nextGoal + INCREMENT : nextGoal;
};
