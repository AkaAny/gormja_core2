
export type Duration=number;

export const MilliSecond:Duration=(1);
export const Second:Duration=1000*MilliSecond;
export const Minute:Duration=60*Second;
export const Hour:Duration=60*Minute;
export const Day:Duration=24*Hour;
export const AverageMonth:Duration=30*Day;