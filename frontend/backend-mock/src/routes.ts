import { Router } from "express";
import { randomArticlesCtrl, organisationsListCtrl } from "./controllers";

export const apiRouter = Router();

apiRouter.get("/articles/random", randomArticlesCtrl);
apiRouter.get("/organisations", organisationsListCtrl);
