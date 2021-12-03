import * as express from "express";
import { apiRouter } from "./routes";

const APP_PORT = 8000;

const app = express();

app.use("/api", apiRouter);

app.listen(APP_PORT, () => {
  console.log(`HTTP Server started on port ${APP_PORT}`);
});
