import { FC } from 'react';

import classes from './Recipe.module.scss';
import { Link } from 'react-router-dom';
import { clientRouter, AppRoutes, apiRouter, ApiRoutes } from '../../../router/router';
import { Heart } from '../../../components/Icon/Heart/Heart';
import { Stars } from '../../../components/Icon/Stars/Stars';

interface Props {
    data: {
        id: number;
        title: string;
        description: string;
        src: string;
        isFavorite: boolean;
        rating: number;
    };
}

export const Recipe: FC<Props> = ({ data }) => {
    const link = clientRouter.buildLink(AppRoutes.recipe, { params: { id: data.id.toString() } });

    return (
        <div className={classes.recipe}>
            <Link to={link}>
                <div className={classes.image_container}>
                    <img src={data.src} alt="" />
                    <p className={classes.description}>{data.description}</p>
                </div>
                <h3>{data.title}</h3>
            </Link>
            <div className={classes.recipe_footer}>
                <Heart />
                <Stars
                    rating={data.rating}
                    onSelected={async (rating) => {
                        const url = apiRouter.buildLink(ApiRoutes.rate_recipe, { params: { id: data.id.toString() } });
                        const response = await fetch(url, {
                            method: 'POST',
                            headers: {
                                'Content-Type': 'application/json',
                            },
                            body: JSON.stringify({
                                rating,
                            }),
                        });

                        const json = await response.json();

                        console.log({ json });
                    }}
                />
            </div>
        </div>
    );
};
