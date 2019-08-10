import React from 'react';
import Drawer from '@material-ui/core/Drawer';
import { MenuItem, List, ListItemIcon, ListItemText, Divider } from '@material-ui/core';
import MoneyOutlined from '@material-ui/icons/MoneyOutlined';
import RestoreOutlined from '@material-ui/icons/RestoreOutlined';
import AccountCircleOutlined from '@material-ui/icons/AccountCircleOutlined';
import { SvgIconProps } from 'material-ui';
import { Link } from "react-router-dom";

interface DrawerItem {
    index: number
    name: string
    link: string
    icon: React.ReactElement<SvgIconProps>
}

const mainDrawerItems: Array<DrawerItem> = [
    {
        index: 0,
        name: 'Calc',
        link: '/calc',
        icon: <MoneyOutlined />,
    },
    {
        index: 1,
        name: 'History',
        link: '/history',
        icon: <RestoreOutlined />,
    },
];

const secondaryDrawerItems: Array<DrawerItem> = [
    {
        index: 2,
        name: 'Account',
        link: '/account',
        icon: <AccountCircleOutlined />,
    },
];

interface State {
    selected: number
}

export default class NavDrawer extends React.Component<{}, State> {
    public state: State = {
        selected: 0
    };

    private setSelected(index: number) {
        this.setState({selected: index})
    }

    public render() {
        return (
            <div>
                <Drawer
                    classes={{paper: 'nav-drawer'}}
                    variant="permanent"
                    anchor="left"
                >
                    <List>
                        {mainDrawerItems.map((item: DrawerItem) => {
                            return (
                                <MenuItem
                                    button
                                    key={item.name}
                                    component={Link}
                                    to={item.link}
                                    onClick={() => this.setSelected(item.index)}
                                    selected={this.state.selected === item.index}
                                >
                                    <ListItemIcon>{item.icon}</ListItemIcon>
                                    <ListItemText primary={item.name} />
                                </MenuItem>
                            );
                        })}
                    </List>
                    <Divider />
                    <List>
                        {secondaryDrawerItems.map((item: DrawerItem) => {
                            return (
                                <MenuItem
                                    button
                                    key={item.name}
                                    component={Link}
                                    to={item.link}
                                    onClick={() => this.setSelected(item.index)}
                                    selected={this.state.selected === item.index}
                                >
                                    <ListItemIcon>{item.icon}</ListItemIcon>
                                    <ListItemText primary={item.name} />
                                </MenuItem>
                            );
                        })}
                    </List>
                </Drawer>
                <main className="content">
                    <div>{this.props.children}</div>
                </main>
            </div>
        );
    }

}
