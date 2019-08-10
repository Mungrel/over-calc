import React from 'react'
import ErrorOutline from '@material-ui/icons/ErrorOutline';

const NotFound: React.FC = () => {
    return (
        <div className="centered-container">
            <div>
                <ErrorOutline />
            </div>
            <div>
                Not Found
            </div>
        </div>
    )
}

export default NotFound
