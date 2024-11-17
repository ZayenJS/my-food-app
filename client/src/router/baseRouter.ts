import { FC } from 'react';
import { ApiRoutes } from './apiRouter';
import { AppRoutes } from './clientRouter';

interface BuildLinkOptions {
    params?: Record<string, string>;
    query?: Record<string, string>;
}

export abstract class BaseRouter<T extends ApiRoutes | AppRoutes> {
    abstract readonly routes: Record<
        T,
        {
            path: string;
            component?: FC;
        }
    >;

    public buildLink(name: T, opts?: BuildLinkOptions): string {
        let link = this.routes[name]?.path;

        if (!link) {
            throw new Error(`No route found for ${name}`);
        }

        if (!opts) {
            return link;
        }

        const params = opts?.params;
        const query = opts?.query;

        if (!params && !query) {
            return link;
        }

        for (const key in params) {
            link = link.replace(`:${key}`, params[key].toString());
        }

        const queryStr = new URLSearchParams(query).toString();

        if (queryStr) {
            link += `?${queryStr}`;
        }

        return link;
    }
}
