import { BigNumber } from 'bignumber.js';

export const shortAccountString = (first: number, last: number, str: string) => {
	return str.substring(0, first) + '...' + str.substring(str.length - last);
};

export const bigIntToFloat = (number: string, decimals: number, precision: number): string => {
	return new BigNumber(number).div(new BigNumber(10).pow(decimals)).toFixed(precision);
};
