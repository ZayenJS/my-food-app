import { FC, PropsWithChildren } from 'react';
import { Header } from '../Header/Header';
import { SideBar } from '../SideBar/SideBar';

import classes from './Layout.module.scss';

export const Layout: FC<PropsWithChildren> = ({ children }) => {
    return (
        <div className={classes.root}>
            <Header />
            <SideBar />
            <main>{children}</main>
        </div>
    );
};
