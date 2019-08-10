import React from 'react';
import { SyncLoader } from 'react-spinners';

const LoadingSpinner: React.FC<{}> = () => {
    return (
        <div className="centered-container">
            <SyncLoader />
        </div>
    );
}
export default LoadingSpinner;
