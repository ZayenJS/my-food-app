import { Burger } from '../Burger/Burger';

import classes from './Header.module.scss';
import logo from '../../assets/images/logo.png';
import { SearchInput } from '../Form/SearchInput/SearchInput';
import { Link } from 'react-router-dom';

export const Header = () => {
    return (
        <header className={classes.header}>
            <div className={classes.header_logo_container}>
                <Burger />
                <Link className={classes.header_logo} to="/">
                    <img className="logo" src={logo} alt="" />
                    <h1>My Food App</h1>
                </Link>
            </div>

            <div>
                <SearchInput placeholder="Rechercher..." />
            </div>
        </header>
    );
};
