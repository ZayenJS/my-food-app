import { FC, FormEvent } from 'react';
import { FoodsIcon } from '../../components/Icon/FoodsIcon/FoodsIcon';

import classes from './NewFood.module.scss';
import { buildClassName } from '../../util';
import { ApiRoutes, apiRouter } from '../../router/apiRouter';

export const NewFood: FC = () => {
    const onSubmitHandler = async (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const method = event.currentTarget.method;
        const url = event.currentTarget.action;
        const formData = new FormData(event.currentTarget);
        const dataAsObject = Object.fromEntries(formData.entries());

        const response = await fetch(url, {
            method,
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(dataAsObject),
        });

        if (!response.ok) {
            console.error('Error while creating the food');
            return;
        }

        const json = await response.json();
        console.log(json);
    };

    return (
        <div className={buildClassName(classes.new_food, 'padded-wrapper-1rem')}>
            <h2 className="page_title">
                <FoodsIcon className="icon-lg" />
                <span>Créer un aliment</span>
            </h2>

            <form action={apiRouter.buildLink(ApiRoutes.foods)} method="POST" onSubmit={onSubmitHandler}>
                <fieldset>
                    <legend>Informations générales</legend>

                    <div className={buildClassName(classes.field_container, classes.general_info_container)}>
                        <div>
                            <label htmlFor="name">Nom</label>
                            <input type="text" id="name" name="name" />
                        </div>
                        <div>
                            <label htmlFor="brand">Marque</label>
                            <input type="text" id="brand" name="brand" />
                        </div>
                        <div>
                            <label htmlFor="category">Catégorie</label>
                            <select id="category" name="category">
                                <option value="breakfast">Petit-déjeuner</option>
                                <option value="lunch">Déjeuner</option>
                                <option value="dinner">Dîner</option>
                                <option value="dessert">Dessert</option>
                                <option value="snack">Collation</option>
                            </select>
                        </div>
                        <div>
                            <label htmlFor="image">Image</label>
                            <input type="file" id="image" name="image" />
                        </div>
                    </div>
                </fieldset>

                <fieldset>
                    <legend>Informations nutritionnelles</legend>

                    <div className={classes.field_container}>
                        <div>
                            <label htmlFor="calories">Calories</label>
                            <input type="number" id="calories" name="calories" />
                        </div>
                        <div>
                            <label htmlFor="proteins">Protéines</label>
                            <input type="number" id="proteins" name="proteins" />
                        </div>
                        <div>
                            <label htmlFor="carbs">Glucides</label>
                            <input type="number" id="carbs" name="carbs" />
                        </div>
                        <div>
                            <label htmlFor="sugar">Sucre</label>
                            <input type="number" id="sugar" name="sugar" />
                        </div>
                        <div>
                            <label htmlFor="fats">Lipides</label>
                            <input type="number" id="fats" name="fats" />
                        </div>
                        <div>
                            <label htmlFor="saturatedFats">Lipides saturés</label>
                            <input type="number" id="saturatedFats" name="saturatedFats" />
                        </div>
                        <div>
                            <label htmlFor="fiber">Fibres</label>
                            <input type="number" id="fiber" name="fiber" />
                        </div>
                        <div>
                            <label htmlFor="salt">Sel</label>
                            <input type="number" id="salt" name="salt" />
                        </div>
                    </div>
                </fieldset>

                <button className="button" type="submit">
                    Créer
                </button>
            </form>
        </div>
    );
};
