import { FC } from 'react';

import classes from './PlusIcon.module.scss';
import { buildClassName } from '../../../util';

interface Props {
    className?: string;
}

export const PlusIcon: FC<Props> = ({ className }) => {
    return (
        <svg
            className={buildClassName(classes.plus_icon, className)}
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none">
            <path
                fillRule="evenodd"
                clipRule="evenodd"
                d="M13 6C13 5.44771 12.5523 5 12 5C11.4477 5 11 5.44771 11 6V11H6C5.44771 11 5 11.4477 5 12C5 12.5523 5.44771 13 6 13H11V18C11 18.5523 11.4477 19 12 19C12.5523 19 13 18.5523 13 18V13H18C18.5523 13 19 12.5523 19 12C19 11.4477 18.5523 11 18 11H13V6Z"
            />
        </svg>
    );
};
