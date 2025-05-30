import express from "express";
import Repository from "./repository/Repository.ts";

const port = 8080;
const tableName = "kvp-table";

const app = express();

app.use(express.json());

const repository = new Repository(tableName);

app.listen(port, () => {
  console.log("Server Listening on port", port);
});

app.get("/:key", async (req, res) => {
  const key = req.params["key"];

  const item = await repository.get(key);

  if (!item) {
    res.status(200).send({ key });
    return;
  }

  res.status(200).send(item);
});

app.put("/:key", async (req, res) => {
  const key = req.params["key"];

  const item = req.body;

  item.key = key;

  const result = await repository.put(item);

  if (!result) {
    res.status(500).end();
    return;
  }

  res.status(200).end();
});

app.patch("/:key", async (req, res) => {
  const key = req.params["key"];

  const item = req.body;

  item.key = key;

  const result = await repository.patch(item);

  if (!result) {
    res.status(500).end();
    return;
  }

  res.status(204).end();
});

app.delete("/:key", (req, res) => {
  const key = req.params["key"];
  const success = repository.delete(key);

  if (!success) {
    res.status(500).end();
    return;
  }
  res.status(204).end();
});
