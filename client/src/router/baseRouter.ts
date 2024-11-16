import { ApiRoutes } from './apiRouter';
import { AppRoutes } from './clientRouter';

export abstract class BaseRouter {
    abstract readonly routes: Record<string, string>;

    public buildLink(name: AppRoutes | ApiRoutes, params?: Record<string, string | number>) {
        let link = this.routes[name];

        if (!params) {
            return link;
        }

        for (const key in params) {
            link = link.replace(`:${key}`, params[key].toString());
        }

        return link;
    }
}
