import { PUBLIC_ROUTER_ADDRESS } from '$env/static/public';
import { Base64 } from '@tonconnect/protocol';
import { beginCell, toNano } from 'ton';

const OP_ADD_BUDGET = 1;
const OP_WITHDRAW_BUDGET = 2;

const addBudgetMessage = async (amount: string) => {
	const tonAmount = toNano(amount);
	const body = beginCell().storeUint(OP_ADD_BUDGET, 32).storeUint(0, 64).endCell();
	const payload = Base64.encode(body.toBoc());

	const msg = {
		address: PUBLIC_ROUTER_ADDRESS,
		amount: tonAmount,
		payload: payload
	};

	return msg;
};

const withdrawBudgetMessage = async (amount: string) => {
	const tonAmount = toNano(amount);

	const body = beginCell()
		.storeUint(OP_WITHDRAW_BUDGET, 32)
		.storeUint(0, 64)
		.storeCoins(tonAmount)
		.endCell();

	const payload = Base64.encode(body.toBoc());

	const msg = {
		address: PUBLIC_ROUTER_ADDRESS,
		payload: payload
	};

	return msg;
};
