import { Link } from 'react-router-dom';

import classes from './SideBar.module.scss';
import { clientRouter } from '../../router/router';

export const SideBar = () => {
    return (
        <aside className={classes.root}>
            <nav className={classes.nav}>
                <Link className={classes.new_recipe} to="/favoris">
                    Favoris
                </Link>
                <hr />

                <ul className={classes.list_top}>
                    <li>
                        <Link to="/recettes/creer">
                            <button>Créer</button>
                        </Link>
                    </li>
                    {clientRouter.sidebarNavigationLinks.map((link) => (
                        <li key={link.id}>
                            <Link to={link.href}>
                                <span>{link.name}</span>
                            </Link>
                        </li>
                    ))}
                </ul>
                <ul className={classes.list_bottom}>
                    <li>
                        <Link to="/parametres">
                            <span>Paramètres</span>
                        </Link>
                    </li>
                </ul>
            </nav>
        </aside>
    );
};
