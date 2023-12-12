CREATE TABLE clinics (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  phone VARCHAR(20) NOT NULL,
  address VARCHAR(255) NOT NULL,
  slot INT NOT NULL,
  created_at BIGINT NOT NULL DEFAULT (CAST (EXTRACT (EPOCH FROM CURRENT_TIMESTAMP) AS BIGINT)),
  updated_at BIGINT NOT NULL DEFAULT (CAST (EXTRACT (EPOCH FROM CURRENT_TIMESTAMP) AS BIGINT))
);