import { fetchApi } from './api';

export interface Operand {
    id: string,
    value: string,
}

interface OperandRequest {
    value: string,
}

const listOperands = (): Promise<Operand[]> => {
    return fetchApi('/operands', {});
}

const createOperand = (operand: OperandRequest): Promise<Operand> => {
    const options = {
        method: 'POST',
        body: JSON.stringify(operand)
    };
    return fetchApi('/operand', options);
}

export const operandService = {
    listOperands,
    createOperand
};
