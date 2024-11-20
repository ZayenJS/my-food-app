import { FC } from 'react';
import { MenuIcon } from '../../components/Icon/MenuIcon/MenuIcon';

export const Menus: FC = () => {
    return (
        <div className="padded-wrapper-1rem">
            <h2 className="page_title">
                <MenuIcon className="icon-lg" />
                <span>Menus</span>
            </h2>
        </div>
    );
};
