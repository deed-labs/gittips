import { Address, beginCell, Cell, Contract, ContractABI, contractAddress, ContractProvider, Sender, SendMode } from 'ton-core';
import fs from 'fs';
import { Maybe } from 'ton-core/dist/utils/maybe';

const budgetJSON = JSON.parse(fs.readFileSync("./build/Budget.compiled.json").toString())

export class Budget  {
    static readonly code = Cell.fromBoc(
        Buffer.from(budgetJSON.hex, "hex")
    )[0];
}
