import React from 'react';
import { Operand } from '../util/operand-service';
import OperandList from '../components/OperandList';
import LoadingSpinner from '../components/LoadingSpinner';
import ErrorLabel from '../components/ErrorLabel';
import { operandService } from '../util/operand-service';

interface State {
    loading: boolean,
    hasError: boolean,
    operands: Operand[],
}

export default class Calc extends React.Component<{}, State> {
    public state: State = {
        loading: true,
        hasError: false,
        operands: [],
    }

    public componentDidMount() {
        operandService.listOperands()
            .then(operands => {
                console.log('operands', operands)
                this.setState({ loading: false, hasError: false, operands: operands});
            }, error => {
                console.log('error', error);
                this.setState({ loading: false, hasError: true, operands: [] });
            })
    }

    public render() {
        const { loading, hasError, operands } = this.state;
        if (loading) {
            return <LoadingSpinner />
        }

        if (hasError) {
            return <ErrorLabel />
        }

        return (
            <OperandList operands={operands}/>
        );
    }
}
