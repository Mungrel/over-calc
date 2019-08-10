import React from 'react';
import { InputLabel } from '@material-ui/core';

interface Props {
    visible: boolean
}

const InvalidLoginDetailsLabel: React.FC<Props> = (props: Props) => {
    return (
        <div className={props.visible ? '' : 'hidden'}>
        <InputLabel error>
          Invalid username or password.
        </InputLabel>
      </div>
    );
};

export default InvalidLoginDetailsLabel;
