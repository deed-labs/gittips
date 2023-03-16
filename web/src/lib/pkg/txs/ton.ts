import { PUBLIC_ROUTER_ADDRESS } from '$env/static/public';
import { Base64 } from '@tonconnect/protocol';
import { beginCell, toNano } from 'ton';

const OP_ADD_BUDGET = 1;
const OP_WITHDRAW_BUDGET = 2;

export interface Message {
	address: string;
	amount: string;
	stateInit?: string | undefined;
	payload?: string | undefined;
}

export const addBudgetMessage = (amount: string): Message => {
	const tonAmount = toNano(amount);
	const body = beginCell().storeUint(OP_ADD_BUDGET, 32).storeUint(0, 64).endCell();
	const payload = Base64.encode(body.toBoc());

	const msg = {
		address: PUBLIC_ROUTER_ADDRESS,
		amount: tonAmount.toString(),
		payload: payload
	};

	return msg;
};

export const withdrawBudgetMessage = (amount: string): Message => {
	const tonAmount = toNano(amount);

	const body = beginCell()
		.storeUint(OP_WITHDRAW_BUDGET, 32)
		.storeUint(0, 64)
		.storeCoins(tonAmount)
		.endCell();

	const payload = Base64.encode(body.toBoc());

	const msg = {
		address: PUBLIC_ROUTER_ADDRESS,
		amount: '0',
		payload: payload
	};

	return msg;
};
