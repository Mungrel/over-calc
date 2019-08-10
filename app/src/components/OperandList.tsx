import React from 'react';
import { Operand } from '../util/operand-service';

interface Props {
    operands: Operand[]
}

const OperandList: React.FC<Props> = (props: Props) => {
    if (props.operands.length === 0) {
        return (
            <div>You have no operands.</div>
        )
    }
    return (
        <ol>
            {props.operands.map(operand => <li>{operand.value}</li>)}
        </ol>
    );
};

export default OperandList;
