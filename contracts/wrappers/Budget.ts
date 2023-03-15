import {
    Address,
    beginCell,
    Cell,
    Contract,
    ContractABI,
    contractAddress,
    ContractProvider,
    Sender,
    SendMode,
} from 'ton-core';
import fs from 'fs';

const budgetJSON = JSON.parse(fs.readFileSync('./build/Budget.compiled.json').toString());

export class Budget implements Contract {
    static readonly code = Cell.fromBoc(Buffer.from(budgetJSON.hex, 'hex'))[0];

    address: Address;

    constructor(address: Address) {
        this.address = address;
    }

    static calculateAddress = (routerAddr: Address, ownerAddr: Address): Address => {
        let data = beginCell().storeAddress(routerAddr).storeAddress(ownerAddr).endCell();
        let init = { code: Budget.code, data };

        return contractAddress(0, init);
    };
}
