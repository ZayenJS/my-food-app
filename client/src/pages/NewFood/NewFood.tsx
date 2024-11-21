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

        const data = {
            ...dataAsObject,
            calories: parseFloat(dataAsObject.calories.toString()),
            protein: parseFloat(dataAsObject.protein.toString()),
            carbs: parseFloat(dataAsObject.carbs.toString()),
            sugar: parseFloat(dataAsObject.sugar.toString()),
            fat: parseFloat(dataAsObject.fat.toString()),
            saturated_fat: parseFloat(dataAsObject.saturated_fat.toString()),
            fiber: parseFloat(dataAsObject.fiber.toString()),
            sodium: parseFloat(dataAsObject.sodium.toString()),
        };

        const response = await fetch(url, {
            method,
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
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
                            <input type="text" id="name" name="name" value="test" />
                        </div>
                        <div>
                            <label htmlFor="brand">Marque</label>
                            <input type="text" id="brand" name="brand" value="Carrefour" />
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
                            <input type="number" id="calories" name="calories" value="127" />
                        </div>
                        <div>
                            <label htmlFor="protein">Protéines</label>
                            <input type="number" id="protein" name="protein" value="20" />
                        </div>
                        <div>
                            <label htmlFor="carbs">Glucides</label>
                            <input type="number" id="carbs" name="carbs" value="65" />
                        </div>
                        <div>
                            <label htmlFor="sugar">Sucre</label>
                            <input type="number" id="sugar" name="sugar" value="11" />
                        </div>
                        <div>
                            <label htmlFor="fat">Lipides</label>
                            <input type="number" id="fat" name="fat" value="5.5" />
                        </div>
                        <div>
                            <label htmlFor="saturated_fat">Lipides saturés</label>
                            <input type="number" id="saturated_fat" name="saturated_fat" value="2.1" />
                        </div>
                        <div>
                            <label htmlFor="fiber">Fibres</label>
                            <input type="number" id="fiber" name="fiber" value="2.1" />
                        </div>
                        <div>
                            <label htmlFor="sodium">Sel</label>
                            <input type="number" id="sodium" name="sodium" value="0.75" />
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
