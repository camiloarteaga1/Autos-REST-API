CREATE TABLE cars (
    id              bigserial                   NOT NULL,
    reservation     bool         DEFAULT false  NULL,
    created_at      timestamptz                 NULL,
    updated_at      timestamptz                 NULL,
    deleted_at      timestamptz                 NULL,
    car_id int8                                 NOT NULL,
    price int8                                  NOT NULL,
    modelYear       int8                        NOT NULL,
    carType         text                        NOT NULL,
    brand           text                        NOT NULL,
    transmission    text                        NOT NULL,
    fuel            text                        NOT NULL,
    motor           text                        NOT NULL,
    doors int8                                  NOT NULL,
    breaks_abs      bool                        NOT NULL,
    price           text                        NOT NULL,
    CONSTRAINT cars_pkey PRIMARY KEY (id)
);

CREATE INDEX idx_cars_deleted_at ON public.cars USING btree (deleted_at);

insert into cars (car_id, reservation, price, modelYear, carType, brand, transmission, fuel, motor, doors, breaks_abs, pic) 
values 
(1, false,  800,    2015,   'Civic',        'Honda',        'Automatic',    'Gasoline',     'Sedan',        4,  true,   'https://www.carpro.com/hs-fs/hubfs/2022-Honda-Civic-feature-carprousa.jpg?width=1530&name=2022-Honda-Civic-feature-carprousa.jpg'),
(2, false,  500,    2018,   'Golf',         'Volkswagen',   'Manual',       'Gasoline',     'Hatchback',    5,  true,   'https://i.gaw.to/vehicles/photos/08/46/084628_2018_volkswagen_Golf.jpg'),
(3, false,  1200,   2022,   'Model 3',      'Tesla',        'Automatic',    'Electric',     'Sedan',        4,  true,   'https://www.mclarencf.com/imagetag/258/5/l/Used-2022-Tesla-Model-3.jpg'),
(4, false,  1000,   2021,   'Tucson',       'Hyundai',      'Automatic',    'Gasoline',     'SUV',          5,  true,   'https://cdn.motor1.com/images/mgl/10Llq/s1/hyundai-tucson-2021-dynamisch.webp'),
(5, false,  600,    2017,   'Fiesta',       'Ford',         'Manual',       'Gasoline',     'Hatchback',    5,  true,   'https://www.ruv.de/dam/jcr:e8284d10-02e6-4b1c-a220-f978505796b5/2017-ford-fiesta-ecoboost-fahrbericht-test-review-kritik-jens-stratmann-1.jpg');
(6, false,  900,    2019,   'Corolla',      'Toyota',       'Automatic',    'Gasoline',     'Sedan',        4,  true,   'https://cdn.motor1.com/images/mgl/K6OEM/s1/new-toyota-corolla-hatchback-hybrid.webp'),
(7, false,  900,    2020,   'Sportage',     'Kia',          'Automatic',    'Hybrid',       'SUV',          5,  true,   'https://www.autocar.co.uk/sites/autocar.co.uk/files/styles/gallery_slide/public/images/car-reviews/first-drives/legacy/kia-sportage-2020-0314.jpg?itok=VE08o5wf'),
(8, false,  1600,   2023,   'Mustang GT',   'Ford',         'Automatic',    'Gasoline',     'Coupe',        2,  true,   'https://www.motortrend.com/uploads/2023/01/2023-Ford-Mustang-Mach-1-6M-21.jpg?fit=around%7C875:492');

