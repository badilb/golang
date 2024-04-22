CREATE TABLE module_info (
                             id BIGSERIAL PRIMARY KEY,
                             created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
                             updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
                             module_name VARCHAR(255) NOT NULL,
                             module_runtime INTEGER NOT NULL,
                             exam_type VARCHAR(255) NOT NULL,
                             version INTEGER NOT NULL DEFAULT 1
);

INSERT INTO module_info (module_name, module_runtime, exam_type, version) VALUES ('Math', 5, '1', 1)