import { FC } from 'react';
import { buildClassName } from '../../../util';
import classes from './MoreVerticalIcon.module.scss';

interface Props {
    className?: string;
}

export const MoreVerticalIcon: FC<Props> = ({ className }) => {
    return (
        <svg
            className={buildClassName(classes.more_vertical_icon, className)}
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none">
            <g id="Menu / More_Vertical">
                <g id="Vector">
                    <path d="M11 18C11 18.5523 11.4477 19 12 19C12.5523 19 13 18.5523 13 18C13 17.4477 12.5523 17 12 17C11.4477 17 11 17.4477 11 18Z" />
                    <path d="M11 12C11 12.5523 11.4477 13 12 13C12.5523 13 13 12.5523 13 12C13 11.4477 12.5523 11 12 11C11.4477 11 11 11.4477 11 12Z" />
                    <path d="M11 6C11 6.55228 11.4477 7 12 7C12.5523 7 13 6.55228 13 6C13 5.44772 12.5523 5 12 5C11.4477 5 11 5.44772 11 6Z" />
                </g>
            </g>
        </svg>
    );
};
