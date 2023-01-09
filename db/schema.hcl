schema "public" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}

table "contents" {
  schema = schema.public
  column "id" {
    null = false
    type = uuid
  }
  column "title_ruby" {
    null = false
    type = text
  }
  column "title" {
    null = false
    type = text
  }
  column "author_ruby" {
    null = false
    type = text
  }
  column "author" {
    null = false
    type = text
  }
  column "speaker_ruby" {
    null = false
    type = text
  }
  column "speaker" {
    null = false
    type = text
  }
  column "file_name" {
    null = false
    type = text
  }
  column "new_arrival_date" {
    null = false
    type = varchar(256)
  }
  column "time" {
    null = false
    type = varchar(256)
  }
  primary_key {
    columns = [column.id]
  }
}
