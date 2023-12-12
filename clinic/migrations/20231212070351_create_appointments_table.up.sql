CREATE TABLE IF NOT EXISTS appointments (
  id SERIAL PRIMARY KEY,
  patient_email VARCHAR(255) NOT NULL REFERENCES patients(email),
  clinic_id INT NOT NULL REFERENCES clinics(id),
  employee_username VARCHAR(50) NOT NULL REFERENCES employees(username),
  appointment_date BIGINT NOT NULL,
  queue_number INT NOT NULL,
  status VARCHAR(20) NOT NULL,
  price INT,
  description VARCHAR(255),
  created_at BIGINT NOT NULL DEFAULT (CAST (EXTRACT (EPOCH FROM CURRENT_TIMESTAMP) AS BIGINT)),
  updated_at BIGINT NOT NULL DEFAULT (CAST (EXTRACT (EPOCH FROM CURRENT_TIMESTAMP) AS BIGINT))
);