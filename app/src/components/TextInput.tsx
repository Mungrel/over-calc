import React from 'react';
import TextField from '@material-ui/core/TextField';

interface Props {
    label: string,
    error: boolean,
    onChange: (value: string) => void
}

const TextInput: React.FC<Props> = (props: Props) => {
    const onValueChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const value = event.target.value;
        props.onChange(value);
    }

    return (
        <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            label={props.label}
            autoComplete="text"
            error={props.error}
            onChange={onValueChange}
            autoFocus
        />
    )
}

export default TextInput;
