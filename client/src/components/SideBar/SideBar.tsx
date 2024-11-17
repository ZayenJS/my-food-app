import { Link, NavLink } from 'react-router-dom';

import classes from './SideBar.module.scss';
import { clientRouter } from '../../router/router';
import { buildClassName } from '../../util';
import { PlusIcon } from '../Icon/PlusIcon/PlusIcon';

export const SideBar = () => {
    return (
        <aside className={classes.sidebar}>
            <nav className={classes.sidebar_nav}>
                <Link className={classes.sidebar_favorite_recipes} to={clientRouter.routes.favorite_recipes.path}>
                    Favoris
                </Link>
                <hr />

                <ul className={classes.sidebar_list_top}>
                    <li>
                        <Link className={classes.sidebar_link} to={clientRouter.routes.recipe_new.path}>
                            <div className={classes.sidebar_new_recipe}>
                                <PlusIcon className={buildClassName(classes.sidebar_link_icon, 'icon-md')} />
                                <span>Créer</span>
                            </div>
                        </Link>
                    </li>
                    {clientRouter.sidebarNavigationLinks.map((link) => (
                        <li key={link.id} className={classes.sidebar_link_container}>
                            <NavLink
                                to={link.href}
                                className={({ isActive }) =>
                                    isActive
                                        ? buildClassName(classes.sidebar_link, classes.active, 'active')
                                        : classes.sidebar_link
                                }>
                                {link.icon ? (
                                    <span className={buildClassName(classes.sidebar_link_icon, 'icon-md')}>
                                        {link.icon}
                                    </span>
                                ) : null}
                                <span>{link.name}</span>
                            </NavLink>
                        </li>
                    ))}
                </ul>
                <ul className={classes.sidebar_list_bottom}>
                    <li>
                        <Link to={clientRouter.routes.settings.path}>
                            <span>Paramètres</span>
                        </Link>
                    </li>
                </ul>
            </nav>
        </aside>
    );
};
