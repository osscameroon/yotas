import { expect } from "chai";
import request from "supertest";
import axios from "axios";
import dotenv from "dotenv";

//Parse env
dotenv.config()
const apiHost = process.env.API_HOST

describe("articles", function () {
  const baseEndpoint = "articles";

  describe("GET", function () {
    it("returns 200 status", async function () {
      return request(apiHost)
        .get(`/${baseEndpoint}`)
        .set("Tenant", "1")
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8");
    });

    it("returns 400 status when Tenant is not set", async function () {
      return request(apiHost)
        .get(`/${baseEndpoint}`)
        .expect(400);
    });
  });

  describe("POST", function () {
    // const baseEndpoint = "articles";

    // it("returns 200 status", async function () {
    //   return request(apiHost)
    //     .post(`/${baseEndpoint}`)
    //     .send({
    //       name: "Oss Stickers",
    //       description: "Stickers for Oss open source projects",
    //       quantity: 10,
    //       price: 5,
    //       pictures: [
    //         {
    //           id: 1,
    //           alt_text: "oss cameroon sticker image",
    //         },
    //       ],
    //       metadata: "string",
    //     })
    //     .set("Tenant", "1")
    //     .expect(200)
    //     .expect("Content-Type", "application/json; charset=utf-8");
    // });

    // it("returns 400 status when Tenant is not set", async function () {
    //   return request(apiHost)
    //     .post(`/${baseEndpoint}`)
    //     .expect(400);
    // });
  });
});

describe("articles/{id}", () => {
  const baseEndpoint = "articles";

  describe("GET", function () {
    it("returns 200 status", async function () {
      return request(apiHost)
        .get(`/${baseEndpoint}/1`)
        .set("Tenant", "1")
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8");
    });

    it("returns 400 status when Tenant is not set", async function () {
      return request(apiHost)
        .get(`/${baseEndpoint}/1`)
        .expect(400);
    });

    it("returns 404 status when the resource is not found", async function () {
      //Please unskip this test once support for this error status is done
      this.skip();
      return request(apiHost)
        .get(`/${baseEndpoint}/404`)
        .set("Tenant", "1")
        .expect(404);
    });
  });

  // describe("DELETE", function () {
  //   const baseEndpoint = "articles";
  //   let article_id;
    // before(function () {
    //   const data = {
    //     name: "data",
    //     description: "data description",
    //     quantity: 10,
    //     price: 2,
    //     pictures: [
    //       {
    //         id: 2,
    //         alt_text: "data picture",
    //       },
    //     ],
    //     metadata: "string",
    //   };
    //
    //   return axios({
    //     url: `${apiHost}/${baseEndpoint}`,
    //     method: "post",
    //     data: data,
    //     headers: { Tenant: "1" },
    //   }).then(function (res) {
    //     article_id = res.data.id;
    //     return res;
    //   });
    // });
    //
    // it("returns 200 status", async function () {
    //   return request(apiHost)
    //     .delete(`/${baseEndpoint}/${article_id}`)
    //     .set("Tenant", "1")
    //     .expect(200)
    //   });

    // it("returns 400 status when Tenant is not set", async function () {
    //   this.skip()
    //   return request(apiHost)
    //     .delete(`/${baseEndpoint}/${article_id}`)
    //     .expect(400);
    // });
    //
    // it("returns 404 status when the resource is not found", async function () {
    //   this.skip()
    //   return request(apiHost)
    //     .delete(`/${baseEndpoint}/404`)
    //     .set("Tenant", "1")
    //     .expect(404);
    // });
  // });

  // describe("PUT", function () {
  //   const baseEndpoint = "articles";
  //   let article_id;
  //   before(function () {
  //     const data = {
  //       name: "data",
  //       description: "data description",
  //       quantity: 10,
  //       price: 2,
  //       pictures: [
  //         {
  //           id: 2,
  //           alt_text: "data picture",
  //         },
  //       ],
  //       metadata: "string",
  //     };
  //
  //     return axios({
  //       url: `${apiHost}/${baseEndpoint}`,
  //       method: "post",
  //       data: data,
  //       headers: { Tenant: "1" },
  //     }).then(function (res) {
  //       article_id = res.data.id;
  //       return res;
  //     });
  //   });
  //
  //   after(function () {
  //     return axios({
  //       url: `${apiHost}/${baseEndpoint}/${article_id}`,
  //       method: "delete",
  //       headers: { Tenant: "1" },
  //     }).then(function (res) {
  //       return res;
  //     });
  //   });
  //
  //   it("returns 200 status", async function () {
  //     const newData = {
  //       name: "new name",
  //       description: "new description",
  //       quantity: -1,
  //       price: -1,
  //       pictures: [
  //         {
  //           id: 1,
  //           alt_text: "oss cameroon sticker image",
  //         },
  //       ],
  //       metadata: "string",
  //     };
  //
  //     return request(apiHost)
  //       .put(`/${baseEndpoint}/${article_id}`)
  //       .send(newData)
  //       .set("Tenant", "1")
  //       .expect(200)
  //       .expect("Content-Type", "application/json; charset=utf-8")
  //       .then((res) => {
  //         expect(res.body.quantity).to.equal(newData.quantity);
  //         expect(res.body.price).to.equal(newData.price);
  //         expect(res.body.metadata).to.equal(newData.metadata);
  //         expect(res.body.name).to.equal(newData.name);
  //         expect(res.body.description).to.equal(newData.description);
  //       })
  //       .catch((err) => {
  //         throw err;
  //       });
  //   });
  //
  //   it("returns 400 status when Tenant is not set", async function () {
  //     return request(apiHost)
  //       .put(`/${baseEndpoint}/1`)
  //       .expect(400);
  //   });
  //
  //   it("returns 404 status when the resource is not found", async function () {
  //     return request(apiHost)
  //       .put(`/${baseEndpoint}/404`)
  //       .set("Tenant", "1")
  //       .expect(404);
  //   });
  // });
});
