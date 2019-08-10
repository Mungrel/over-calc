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
import TextInput from '../components/TextInput';
import ErrorLabel from '../components/ErrorLabel';

interface State {
  username: string;
  password: string;
  error: string | null;
}

export default class SignUp extends React.Component<{}, State> {
  public state: State = {
    username: '',
    password: '',
    error: null,
  }

  private onUsernameChange = (value: string) => {
    this.setState({ username: value });
  }

  private onPasswordChange = (value: string) => {
    this.setState({ password: value });
  }

  private onSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    const { username, password } = this.state;
    console.log('username', username);
    userService.signUp(username, password)
      .then(user => {
        console.log('user', user)
        this.setState( { error: null })
        window.location.href = '/calc'
      }, err => {
        console.log('error', err);
        this.setState({ error: err })
      })
  }

  public render() {
    if (this.state.error) {
      return (
        <ErrorLabel />
      )
    }

    return (
      <SignUpForm
        onUsernameChange={this.onUsernameChange}
        onPasswordChange={this.onPasswordChange}
        onSubmit={this.onSubmit}
        error={this.state.error}
      />
    )
  }
}

interface Props {
  onPasswordChange: (value: string) => void;
  onUsernameChange: (value: string) => void;
  onSubmit: (event: React.FormEvent) => void;
  error: string | null;
}

const SignUpForm: React.FC<Props> = (props: Props) => {
  const classes = useStyles();
  return (
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      <div className={classes.paper}>
        <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Sign up
        </Typography>
        <form
          className={classes.form}
          noValidate
          onSubmit={props.onSubmit}
        >
          <Grid container spacing={2}>
            <Grid item xs={12}>
              <TextInput
                label="Username"
                error={!!props.error}
                onChange={props.onUsernameChange}
              />
            </Grid>
            <Grid item xs={12}>
              <TextInput
                label="Password"
                error={!!props.error}
                onChange={props.onPasswordChange}
              />
            </Grid>
          </Grid>
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
            onSubmit={props.onSubmit}
          >
            Sign Up
          </Button>
          <Grid container justify="flex-end">
            <Grid item>
              <Link href="/sign_in" variant="body2">
                Already have an account? Sign in
              </Link>
            </Grid>
          </Grid>
        </form>
      </div>
    </Container>
  );
}
