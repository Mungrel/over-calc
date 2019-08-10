import React from 'react';
import ErrorOutline from '@material-ui/icons/ErrorOutline';
import { Button } from '@material-ui/core';

const ErrorLabel: React.FC<{}> = () => {
    return (
        <div className="centered-container">
            <div>
                <ErrorOutline />
            </div>
            <div>
                Looks like something's not quite right.
            </div>
            <div>
                <Button onClick={refresh}>Retry</Button>
            </div>
        </div>
    );
}

const refresh = () => {
    console.log('refreshing...');
    window.location.reload();
}

export default ErrorLabel;
