export const shortAccountString = (first: number, last: number, str: string) => {
	return str.substring(0, first) + '...' + str.substring(str.length - last);
};
