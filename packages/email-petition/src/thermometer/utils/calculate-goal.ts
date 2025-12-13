const MIN_GOAL = 20000;
const INCREMENT = 10000;
const THRESHOLD = 1000;

export const getNextGoal = (amt: number) => {
  if (amt < MIN_GOAL - THRESHOLD) {
    return MIN_GOAL;
  }
  for (let goal = MIN_GOAL; ; goal += INCREMENT) {
    if (goal - amt > THRESHOLD) {
      return goal;
    }
  }
};
