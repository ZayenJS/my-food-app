import { FC } from 'react';

import classes from './Burger.module.scss';

export const Burger: FC = () => {
    return (
        <div className={classes.burger}>
            <div className={classes.burger_line} />
            <div className={classes.burger_line} />
            <div className={classes.burger_line} />
        </div>
    );
};
