import { FC } from 'react';

import classes from './Burger.module.scss';

export const Burger: FC = () => {
    return (
        <div className={classes.root}>
            <div className={classes.line} />
            <div className={classes.line} />
            <div className={classes.line} />
        </div>
    );
};
