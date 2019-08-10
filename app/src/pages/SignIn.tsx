import React from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import Container from '@material-ui/core/Container';
import { userService } from '../util/user-service';
import useStyles from '../util/use-styles';
import InvalidLoginDetailsLabel from '../components/InvalidLoginDetailsLabel';
import TextInput from '../components/TextInput';

interface State {
  username: string,
  password: string,
  invalidUserOrPass: boolean,
}

export default class SignIn extends React.Component<{}, State>{
  public state: State = {
    username: '',
    password: '',
    invalidUserOrPass: false,
  };


  private onSubmit = (event: any): void => {
    event.preventDefault();
    const { username, password } = this.state;
    userService.login(username, password)
      .then(user => {
        console.log('user', user);
        window.location.href = '/calc';
      }, error => {
        console.log('error', error);
        this.setState({invalidUserOrPass: true})
      });
  }

  private onPasswordChange = (value: string) => {
    this.setState({ password: value });
  }

  private onUsernameChange = (value: string) => {
    this.setState({ username: value });
  }

  public render() {
    return (
      <SignInForm
        onPasswordChange={this.onPasswordChange}
        onUsernameChange={this.onUsernameChange}
        onSubmit={this.onSubmit}
        invalidUserOrPass={this.state.invalidUserOrPass}
      />
    )
  }
};

interface Props {
  onPasswordChange: (value: string) => void;
  onUsernameChange: (value: string) => void;
  onSubmit: (event: React.FormEvent) => void;
  invalidUserOrPass: boolean;
}

const SignInForm: React.FC<Props> = (props: Props) => {
  const classes = useStyles({});
  console.log('invalidUserOrPass', props.invalidUserOrPass);
  return (
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      <div className={classes.paper}>
        <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Sign in
        </Typography>
        <form
          className={classes.form}
          onSubmit={props.onSubmit}
          noValidate
        >
            <TextInput
                label="Username"
                error={props.invalidUserOrPass}
                onChange={props.onUsernameChange}
            />
            <TextInput
                label="Password"
                error={props.invalidUserOrPass}
                onChange={props.onPasswordChange}
            />
            <InvalidLoginDetailsLabel visible={props.invalidUserOrPass} />
            <Button
                type="submit"
                fullWidth
                variant="contained"
                color="primary"
                className={classes.submit}
            >
                Sign In
            </Button>
            <Grid container>
                <Grid item>
                    <Link href="/sign_up" variant="body2">
                        {"Don't have an account? Sign Up"}
                    </Link>
                </Grid>
            </Grid>
        </form>
      </div>
    </Container>
  );
}
