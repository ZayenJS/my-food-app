import { FC, useRef } from 'react';
import { Star } from '../Star/Star';
import { Random } from '../../../util/random';

import classes from './Stars.module.scss';

interface Props {
    rating: number;
    onSelected?: (rating: number) => void;
}

export const Stars: FC<Props> = ({ rating, onSelected }) => {
    const containerRef = useRef<HTMLDivElement>(null);
    const internalRating = useRef(rating);

    const handleMouseEnter = (index: number) => {
        if (!containerRef.current) {
            return;
        }

        const stars = containerRef.current.children;
        const starsArray = Array.from(stars);

        starsArray.forEach((star, i) => {
            if (i + 1 <= index) {
                return star.classList.add('filled');
            }

            star.classList.remove('filled');
        });
    };

    const handleMouseLeave = () => {
        if (!containerRef.current) {
            return;
        }

        const stars = containerRef.current.children;
        const starsArray = Array.from(stars);

        starsArray.forEach((star, i) => {
            star.classList.remove('filled');

            if (i + 1 <= internalRating.current) {
                star.classList.add('filled');
            }
        });
    };

    const handleClick = async (index: number) => {
        internalRating.current = index;
        onSelected?.(index);
    };

    return (
        <div role="presentation" ref={containerRef} onMouseLeave={handleMouseLeave} className={classes.root}>
            {Array.from({ length: 5 }).map((_, index) => {
                const isFilled = index < internalRating.current;

                return (
                    <Star
                        key={Random.buildRandomId()}
                        index={index}
                        isFilled={isFilled}
                        onMouseEnter={handleMouseEnter}
                        onClick={handleClick}
                    />
                );
            })}
        </div>
    );
};
