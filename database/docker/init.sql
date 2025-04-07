CREATE TABLE IF NOT EXISTS gyms (
    gym_id SERIAL PRIMARY KEY,
    gym_name VARCHAR(255) NOT NULL,
    gym_type VARCHAR(50) NOT NULL,
    gym_office_hours  VARCHAR(50) NOT NULL,
    gym_location TEXT NOT NULL,
    gym_contact_number VARCHAR(50) NOT NULL,
    gym_email VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS gym_classes (
    class_id SERIAL PRIMARY KEY,
    class_name VARCHAR(255) NOT NULL,
    description text NOT NULL,
    class_type VARCHAR(50) NOT NULL,
    class_duration INT NOT NULL,
    class_schedule VARCHAR(255) NOT NULL,
    class_price DECIMAL(10, 2) NOT NULL,
    class_level VARCHAR(50) NOT NULL,
    class_url TEXT NOT NULL,
    gym_id INT NOT NULL,
    FOREIGN KEY (gym_id) REFERENCES gyms(gym_id)
);

CREATE TABLE IF NOT EXISTS affiliators (
    affiliator_id SERIAL PRIMARY KEY,
    affiliator_fname VARCHAR(255) NOT NULL,
    affiliator_lname VARCHAR(255) NOT NULL,
    affiliator_email VARCHAR(255) UNIQUE NOT NULL,
    affiliator_phone VARCHAR(50) NOT NULL,
    api_key TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS affiliator_classses (
    affiliator_id INT NOT NULL,
    gym_class_id INT NOT NULL,
    PRIMARY KEY (affiliator_id, gym_class_id),
    FOREIGN KEY (affiliator_id) REFERENCES affiliators(affiliator_id),
    FOREIGN KEY (gym_class_id) REFERENCES gym_classes(class_id)
);

CREATE TABLE IF NOT EXISTS referrer_websites (
    referrer_id SERIAL PRIMARY KEY,
    affiliator_id INT NOT NULL,
    website_url TEXT,
    FOREIGN KEY (affiliator_id) REFERENCES affiliators(affiliator_id)
);

CREATE TABLE IF NOT EXISTS log_api_requests (
    request_id SERIAL PRIMARY KEY,
    affiliator_id INT NOT NULL,
    affiliator_fname VARCHAR(255) NOT NULL,
    affiliator_lname VARCHAR(255) NOT NULL,
    request_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    endpoint TEXT NOT NULL,
    path_parameters TEXT NULL,
    query_parameters TEXT NULL,
    method TEXT NOT NULL,
    FOREIGN KEY (affiliator_id) REFERENCES affiliators(affiliator_id)
);

CREATE TABLE IF NOT EXISTS click_logs (
    click_id SERIAL PRIMARY KEY,
    class_url TEXT NOT NULL,
    class_id INT NOT NULL,
    class_type VARCHAR(50) NOT NULL,
    click_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMIT;

INSERT INTO gyms (gym_name, gym_type, gym_office_hours, gym_location, gym_contact_number, gym_email) VALUES 
('FitZone Gym', 'ฟิตเนสและเวทเทรนนิ่ง', '06:00 - 22:00', '123 Sukhumvit Road, Bangkok', '098-123-4567', 'contact@fitzone.com'),
('Elite Boxing Club', 'มวยไทยและศิลปะการต่อสู้', '07:00 - 20:00', '45 Silom Road, Bangkok', '089-987-6543', 'info@eliteboxing.com'),
('Flex Yoga Studio', 'โยคะและพิลาทิส', '08:00 - 21:00', '99 Rama IV Road, Bangkok', '092-345-6789', 'hello@flexyoga.com'),
('AquaFit Center', 'กีฬาทางน้ำ', '09:00 - 18:00', '777 Charoenkrung Road, Bangkok', '096-789-4321', 'support@aquafit.com'),
('Groove Dance Studio', 'Dance Fitness และ Zumba', '10:00 - 21:00', '55 Sukhumvit Soi 11, Bangkok', '097-222-3456', 'contact@groovedance.com');

INSERT INTO gym_classes (class_name, description, class_type, class_duration, class_schedule, class_price, class_level, class_url, gym_id) VALUES 
('Reformer Pilates - Foundations', 
 'เป็นคลาสที่เพอร์เฟ็กต์สำหรับผู้เริ่มต้นฝึกพิลาทิส ผู้ฝึกจะได้รับการดูแลและคำแนะนำจากครูผู้สอน โดยคลาสนี้จะพาผู้ฝึกไปรู้จักกับพิลาทิส รีฟอร์เมอร์ ตั้งแต่การใช้เครื่อง ไปจนถึงท่าออกกำลังต่างๆ ที่สามารถนำไปต่อยอดในการฝึกคลาสพิลาทิสอื่นๆได้', 
 'Reformer Pilates', 45, '2025-05-03', 600, 'ง่าย', 'https://waenwonga.wixsite.com/gym3/class1', (SELECT gym_id FROM gyms WHERE gym_name = 'Flex Yoga Studio' LIMIT 1)),
('Power Yoga', 
 'พาวเวอร์โยคะคลาสนี้ไม่ได้ท้าทายแค่ความ "เซน" ภายในตัวของคุณ แต่ยังท้าทายความแข็งแกร่งของคุณอีกด้วย คุณจะได้เคลื่อนไหวแบบต่อเนื่อง ด้วยท่วงท่าที่จะช่วยเสริมสร้างความยืดหยุ่น แต่ก็ยังช่วยเพิ่มความแข็งแกร่งของแกนกลางลำตัว และความอดทนของกล้ามเนื้อที่เกี่ยวข้อง', 
 'Yoga', 60, '2025-05-07', 500, 'ปานกลาง', 'https://waenwonga.wixsite.com/gym3/class2', (SELECT gym_id FROM gyms WHERE gym_name = 'Flex Yoga Studio' LIMIT 1)),
('BODYBALANCE', 
 'เป็นคลาสโยคะแบบใหม่ ที่เหมาะกับทุกคน ด้วยส่วนผสมที่ลงตัวระหว่างโยคะ ไทชิ และพิลาทีส เพื่อเพิ่มความยืดหยุ่นและความแข็งแรง เสริมพื้นฐานที่สำคัญสำหรับร่างกายในการออกกำลังกายในรูปแบบอื่น ๆ', 
 'Fusion Yoga', 60, '2025-05-08', 550, 'ปานกลาง', 'https://waenwonga.wixsite.com/gym3/class3', (SELECT gym_id FROM gyms WHERE gym_name = 'Flex Yoga Studio' LIMIT 1)),
('Pilates with Props', 
 'คลาสพิลาทิสที่ใช้ลูกบอล โฟมโรลเลอร์ หรือวงแหวนเพื่อเพิ่มความท้าทายและช่วยเสริมสร้างความแข็งแรงของกล้ามเนื้อ', 
 'Equipment Pilates', 45, '2025-05-17', 650, 'ปานกลาง', 'https://waenwonga.wixsite.com/gym3/class4', (SELECT gym_id FROM gyms WHERE gym_name = 'Flex Yoga Studio' LIMIT 1)),
('Hot Yoga', 
 'คลาสโยคะที่ฝึกในห้องที่มีอุณหภูมิสูง ช่วยกระตุ้นการไหลเวียนโลหิต เพิ่มความยืดหยุ่นของกล้ามเนื้อ และช่วยขับสารพิษออกจากร่างกาย', 
 'Hot Yoga', 60, '2025-05-13', 600, 'ปานกลาง', 'https://waenwonga.wixsite.com/gym3/class5', (SELECT gym_id FROM gyms WHERE gym_name = 'Flex Yoga Studio' LIMIT 1)),
('Yoga Flow', 
 'โยคะโฟลว์ เป็นการเคลื่อนไหวแบบวินยาสะโยคะ ครูจะพาให้คุณเคลื่อนไหวร่างกายอย่างไหลลื่นไปพร้อมกับการกำหนดลมหายใจ', 
 'Yoga', 60, '2025-05-11', 520, 'ปานกลาง', 'https://waenwonga.wixsite.com/gym3/class6', (SELECT gym_id FROM gyms WHERE gym_name = 'Flex Yoga Studio' LIMIT 1));


INSERT INTO gym_classes (class_name, description, class_type, class_duration, class_schedule, class_price, class_level, class_url, gym_id) VALUES 
('Body Pump', 
 'ยกน้ำหนักที่ผสมผสานกับการเคลื่อนไหวตามจังหวะเพลง เสริมสร้างกล้ามเนื้อและความแข็งแรงทั่วร่างกาย', 
 'Strength and Conditioning', 60, '2025-05-03', 300, 'ปานกลาง', 'https://waenwonga.wixsite.com/gym1/class2', (SELECT gym_id FROM gyms WHERE gym_name = 'FitZone Gym' LIMIT 1)),
('Core Abs', 
 'เน้นการเสริมสร้างกล้ามเนื้อหน้าท้องและแกนกลางของร่างกาย', 
 'Strength and Conditioning', 45, '2025-05-06', 200, 'ปานกลาง', 'https://waenwonga.wixsite.com/gym1/class3', (SELECT gym_id FROM gyms WHERE gym_name = 'FitZone Gym' LIMIT 1)),
('Group Suspension Exercise Class (TRX)', 
 'ใช้สาย TRX ในการฝึกซ้อมเสริมสร้างกล้ามเนื้อและความสมดุลของร่างกาย', 
 'Strength and Conditioning', 45, '2025-05-10', 350, 'ปานกลาง', 'https://waenwonga.wixsite.com/gym1/class4', (SELECT gym_id FROM gyms WHERE gym_name = 'FitZone Gym' LIMIT 1)),
('Body Combat', 
 'คาร์ดิโอที่ผสมผสานการเคลื่อนไหวจากศิลปะการต่อสู้ เช่น คาราเต้ มวยไทย และเทควันโด', 
 'Cardio', 60, '2025-05-12', 300, 'ปานกลางถึงสูง', 'https://waenwonga.wixsite.com/gym1/class5', (SELECT gym_id FROM gyms WHERE gym_name = 'FitZone Gym' LIMIT 1)),
('RPM (Cycling)', 
 'ปั่นจักรยานในร่ม เน้นการเผาผลาญแคลอรี่และเสริมสร้างความแข็งแรงของขา', 
 'Cycling', 45, '2025-05-14', 250, 'ปานกลางถึงสูง', 'https://waenwonga.wixsite.com/gym1/class6', (SELECT gym_id FROM gyms WHERE gym_name = 'FitZone Gym' LIMIT 1)),
('Bootcamp', 
 'การออกกำลังกายที่ผสมผสานหลายประเภท เพื่อเพิ่มความแข็งแรงและความทนทาน', 
 'Strength/Cardio', 60, '2025-05-19', 300, 'สูง', 'https://waenwonga.wixsite.com/gym1/class1', (SELECT gym_id FROM gyms WHERE gym_name = 'FitZone Gym' LIMIT 1));

INSERT INTO gym_classes (class_name, description, class_type, class_duration, class_schedule, class_price, class_level, class_url, gym_id) VALUES 
('HIIT Boom', 
'คลาส HIIT Boom เป็นการออกกำลังกายที่มีความเข้มข้นสูง (HIIT) ซึ่งได้รับแรงบันดาลใจจากการชกมวย เพื่อท้าทายความอดทนและพัฒนากล้ามเนื้อในหลายส่วนของร่างกาย ทั้งกล้ามเนื้อแขน, กล้ามหน้าท้อง, และกล้ามเนื้อส่วนบนและส่วนล่าง คลาสนี้จะช่วยเพิ่มความแข็งแกร่งและความยืดหยุ่นในทุกท่าทาง ด้วยเทคนิคการชกมวยพื้นฐาน เช่น แยป (Jab), ครอส (Cross), ฮุค (Hook), ฟุตเวิร์ค (Footwork) และอื่น ๆ ที่จะทำให้คุณรู้สึกตื่นเต้นและเผาผลาญแคลอรี่ได้อย่างรวดเร็ว', 
'การออกกำลังกาย', 60, '2025-05-01', 450, 'สูง', 'https://waenwonga.wixsite.com/gym2/class1', (SELECT gym_id FROM gyms WHERE gym_name = 'Elite Boxing Club' LIMIT 1)),
('มวยไฟท์โปร', 
'คลาส มวยไฟท์โปร เป็นการออกกำลังกายแบบฟรีสไตล์ที่ได้รับแรงบันดาลใจจากมวยไทย โดยการฝึกฝนจะผสมผสานการใช้กระสอบทรายเข้ากับท่าออกกำลังกายบนพื้น เพื่อเสริมสร้างความแข็งแกร่งและความคล่องตัวให้กับร่างกาย คลาสนี้เน้นการเคลื่อนไหวอย่างอิสระในทุกทิศทาง โดยใช้เพียงน้ำหนักตัวของคุณในการฝึกฝน ช่วยพัฒนากล้ามเนื้อและเพิ่มความทนทานในแต่ละส่วนของร่างกาย', 
'การออกกำลังกาย', 60, '2025-05-07', 500, 'ปานกลางถึงสูง', 'https://waenwonga.wixsite.com/gym2/class2', (SELECT gym_id FROM gyms WHERE gym_name = 'Elite Boxing Club' LIMIT 1)),
('Street Fight Defense', 
'คลาส Street Fight Defense ออกแบบมาเพื่อสอน เทคนิคการป้องกันตัวในสถานการณ์จริง ที่อาจเกิดขึ้นในชีวิตประจำวัน เช่น การถูกจู่โจม, การป้องกันการล็อกตัว, หรือการเอาตัวรอดจากการถูกโจมตีโดยไม่คาดคิด (Self-Defense Training)', 
'ศิลปะป้องกันตัว', 75, '2025-05-03', 400, 'ปานกลาง', 'https://waenwonga.wixsite.com/gym2/class3', (SELECT gym_id FROM gyms WHERE gym_name = 'Elite Boxing Club' LIMIT 1)),
('MMA Grappling', 
'คลาส MMA Grappling (MMA การจับล็อก) ออกแบบมาเพื่อฝึกฝนทักษะการจับล็อกและการทุ่มในศิลปะการต่อสู้แบบผสม (MMA) โดยเน้นที่เทคนิคการจับคู่ต่อสู้ลงพื้น การควบคุมคู่ต่อสู้ และการใช้ท่าทุ่มอย่างมีประสิทธิภาพในสถานการณ์การต่อสู้จริง คลาสนี้เหมาะสำหรับผู้ที่ต้องการพัฒนาทักษะในการจัดการกับคู่ต่อสู้ในระยะใกล้และการต่อสู้บนพื้น (Ground Fighting) ผู้เรียนจะได้เรียนรู้การใช้ท่าทางและกลยุทธ์ต่างๆ ในการจับล็อก การยืดอำนาจ และการหลบหลีกท่าทุ่มของคู่ต่อสู้ พร้อมได้รับคำแนะนำจากโค้ชที่มีประสบการณ์', 
'ศิลปะป้องกันตัว', 120, '2025-05-22', 550, 'สูง', 'https://waenwonga.wixsite.com/gym2/class4', (SELECT gym_id FROM gyms WHERE gym_name = 'Elite Boxing Club' LIMIT 1)),
('Muay Thai Sparring', 
'คลาส Muay Thai Sparring เป็นหลักสูตรที่ออกแบบมาเพื่อผู้ที่มีพื้นฐานมวยไทยและต้องการฝึกซ้อมแบบลงนวมกับคู่ซ้อมจริง คลาสนี้เน้นการฝึกทักษะมวยไทยในสถานการณ์จริง เพื่อพัฒนาความแม่นยำในการโจมตี การหลบหลีก และการป้องกันตัว รวมถึงการเรียนรู้การรับมือกับคู่ต่อสู้ในสภาพแวดล้อมที่มีการปฏิสัมพันธ์โดยตรง เช่น การใช้หมัด, เข่า, ศอก และเท้าในการโจมตี การฝึกในคลาสนี้จะช่วยให้คุณพัฒนาความมั่นใจและเตรียมตัวสำหรับการแข่งขันหรือการใช้งานในสถานการณ์จริง', 
'ศิลปะป้องกันตัว และมวยไทย', 120, '2025-05-12', 500, 'สูง', 'https://waenwonga.wixsite.com/gym2/class5', (SELECT gym_id FROM gyms WHERE gym_name = 'Elite Boxing Club' LIMIT 1)),
('Brazilian Jiu-Jitsu (BJJ) Basics', 
'คลาส Brazilian Jiu-Jitsu (BJJ) Basics เป็นหลักสูตรเบื้องต้นสำหรับผู้ที่สนใจเรียนรู้ทักษะ BJJ ซึ่งเน้นการควบคุมคู่ต่อสู้ผ่านท่าล็อกและท่าทุ่มบนพื้น คลาสนี้จะสอนพื้นฐานของการจับล็อก การควบคุมท่าต่างๆ รวมถึงเทคนิคการป้องกันตัวจากการถูกจับล็อกในท่าที่ไม่คุ้นเคย คลาสนี้เหมาะสำหรับผู้เริ่มต้นที่ต้องการทำความเข้าใจพื้นฐานการต่อสู้แบบ BJJ และการใช้ท่าล็อกในการจัดการกับคู่ต่อสู้ในสถานการณ์ต่างๆ โดยมีโค้ชที่มีประสบการณ์คอยแนะนำและสอนเทคนิคที่ถูกต้อง', 
'ศิลปะป้องกันตัว', 90, '2025-05-25', 450, 'ปานกลาง', 'https://waenwonga.wixsite.com/gym2/class6', (SELECT gym_id FROM gyms WHERE gym_name = 'Elite Boxing Club' LIMIT 1));

-- INSERT INTO gym_classes (class_name, description, class_type, class_duration, class_schedule, class_price, class_level, gym_id) 
-- VALUES 
-- ('H2 Flow', 
--  'คลาสนี้ผสมผสานการเคลื่อนไหวแบบช้าและมีสมาธิ เพื่อเสริมสร้างความยืดหยุ่น และความแข็งแรงให้กับร่างกาย ช่วยเพิ่มความสงบและสมดุลให้กับจิตใจ', 
--  'Mind and Body', 60, '2025-05-01', 350, 'ปานกลาง', (SELECT gym_id FROM gyms WHERE gym_name = 'AquaFit Center' LIMIT 1)),
-- ('Aqua Fit', 
--  'คลาสนี้ใช้การออกกำลังกายในน้ำเพื่อเพิ่มความแข็งแรงและความทนทานของร่างกาย โดยจะช่วยในการลดความเสี่ยงจากการบาดเจ็บและช่วยเผาผลาญแคลอรีได้ดี', 
--  'Strength and Conditioning', 45, '2025-05-03', 400, 'ปานกลาง', (SELECT gym_id FROM gyms WHERE gym_name = 'AquaFit Center' LIMIT 1)),
-- ('H2 Core', 
--  'คลาสที่เน้นการฝึกกล้ามเนื้อแกนกลางร่างกาย (Core) ด้วยการออกกำลังกายในน้ำ เพิ่มความแข็งแรงและความมั่นคงให้กับลำตัว', 
--  'Strength and Conditioning', 45, '2025-05-05', 380, 'ปานกลางถึงสูง', (SELECT gym_id FROM gyms WHERE gym_name = 'AquaFit Center' LIMIT 1)),
-- ('Aqua Splash', 
--  'คลาสที่สนุกสนานและเป็นการออกกำลังกายในน้ำช่วยเพิ่มความแข็งแรงให้กับร่างกายและช่วยเสริมสร้างความคล่องตัว', 
--  'Aqua Fitness', 60, '2025-05-15', 420, 'ปานกลาง', (SELECT gym_id FROM gyms WHERE gym_name = 'AquaFit Center' LIMIT 1)),
-- ('Aqua Bootcamp', 
--  'คลาสที่ออกแบบมาเพื่อฝึกทักษะการออกกำลังกายอย่างเข้มข้นในน้ำ ช่วยเพิ่มความแข็งแรงและการเผาผลาญแคลอรี', 
--  'Aqua Fitness', 60, '2025-05-30', 450, 'สูง', (SELECT gym_id FROM gyms WHERE gym_name = 'AquaFit Center' LIMIT 1)),
-- ('Dance Aqua', 
--  'คลาสเต้นในน้ำที่สนุกสนาน พร้อมการออกกำลังกายแบบคาร์ดิโอ ช่วยเสริมสร้างความแข็งแรงและความยืดหยุ่น', 
--  'Aqua Fitness', 45, '2025-06-05', 400, 'ปานกลาง', (SELECT gym_id FROM gyms WHERE gym_name = 'AquaFit Center' LIMIT 1));

-- INSERT INTO gym_classes (class_name, description, class_type, class_duration, class_schedule, class_price, class_level, gym_id) VALUES 
-- ('Zumba', 
--  'เป็นการออกกำลังโดยการเต้นท่วงท่าที่ใช้พลังน้อย กับใช้พลังเยอะสลับกันไป ประกอบเข้ากับเพลงที่แสนสนุก และท่าเต้นสไตล์ เมเรงเก้ ซัลซ่า เร็กเก้ และคัมเบีย ที่จะทำให้คุณก้าวขาออกไปจากห้องนี้ด้วยพลังที่เต็มเปี่ยม คุณไม่จำเป็นต้องมีประสบการณ์ด้านการเต้นมาก่อน เพราะเราจะช่วยให้คุณได้เสียเหงื่อด้วยความสนุกของคลาสนี้เอง!', 
--  'Dance Fitness', 45, '2025-05-03', 300, 'ง่าย', (SELECT gym_id FROM gyms WHERE gym_name = 'Groove Dance Studio' LIMIT 1)),
-- ('Les Mills BODYJAM: Hip-Hop Groove', 
--  'คลาสเต้นที่รวมเอาจังหวะฮิปฮอปเข้ากับการออกกำลังกายแบบเต็มตัว การเต้นจะเน้นการเคลื่อนไหวที่มีพลังและการสื่อสารผ่านการเต้น', 
--  'Dance Fitness, Street Dance', 50, '2025-05-04', 300, 'ปานกลาง', (SELECT gym_id FROM gyms WHERE gym_name = 'Groove Dance Studio' LIMIT 1)),
-- ('K-Pop Dance Party', 
--  'ถ้าคุณรักในจังหวะของ K-Pop มาร่วมสนุกกับคลาส K-Pop Dance Party ที่จะทำให้คุณได้เต้นตามจังหวะดนตรีที่มีความสนุกสนานและท่าเต้นที่สร้างพลัง คลาสนี้จะพาคุณไปสัมผัสประสบการณ์การเต้นที่เต็มไปด้วยความสนุก พร้อมทั้งเพิ่มความแข็งแรงและความยืดหยุ่นให้กับร่างกาย', 
--  'Dance Fitness, Pop Dance', 50, '2025-05-11', 280, 'ปานกลาง', (SELECT gym_id FROM gyms WHERE gym_name = 'Groove Dance Studio' LIMIT 1)),
-- ('Tango Passion', 
--  'Tango Passion คือการเต้นที่มีความร้อนแรงและอารมณ์สูง เน้นท่าเต้นที่ละเอียดและการจับคู่ที่เข้ากันได้ดี เต้นไปกับท่วงท่าที่มีพลังและความเป็นศิลปะ คลาสนี้จะช่วยให้คุณได้สัมผัสกับความงามของการเต้นท่าทางที่มีเสน่ห์และเพิ่มความแข็งแรงให้กับร่างกาย', 
--  'Latin Dance', 60, '2025-05-17', 400, 'สูง', (SELECT gym_id FROM gyms WHERE gym_name = 'Groove Dance Studio' LIMIT 1)),
-- ('Street Jazz Vibes', 
--  'Street Jazz เป็นการเต้นที่ผสมผสานระหว่าง Jazz และ Hip-Hop การเคลื่อนไหวจะมีความยืดหยุ่นและทันสมัยมากขึ้น โดยมีการผสมกลิ่นของ Hip-Hop เข้าไปในท่าทาง ซึ่งทำให้มีความสนุกสนานและหลากหลาย ไม่จำเป็นต้องมีพื้นฐานในการเต้นมาก่อน เริ่มต้นที่ Basic หรือ Beginner', 
--  'Jazzercise, Dance Fitness', 50, '2025-05-02', 280, 'ง่าย', (SELECT gym_id FROM gyms WHERE gym_name = 'Groove Dance Studio' LIMIT 1)),
-- ('Les Mills BODYJAM', 
--  'เป็นคลาสเต้นเพื่อการออกกำลังกายที่มาพร้อมกับเพลงและท่าเต้นที่มีพลัง ในแต่ละคลาสจะประกอบด้วย 10 เพลงที่ออกแบบมาเพื่อให้คุณได้ปลดปล่อยพลังในท่าทางที่มีความหลากหลาย ไม่ว่าจะเป็น Hip-Hop, Dancehall, House หรือแม้กระทั่งการเต้นที่มีการผสมผสานของสไตล์ต่าง ๆ คุณจะได้ทั้งวอร์มอัพ, บริหารกล้ามเนื้อหัวใจ, และการคูลดาวน์ภายใน 60 นาที คลาสนี้เหมาะสำหรับทุกระดับและจะทำให้คุณรู้สึกสนุกไปกับการออกกำลังกายที่ไม่ซ้ำซาก', 
--  'Dance Fitness', 60, '2025-05-10', 350, 'ปานกลาง', (SELECT gym_id FROM gyms WHERE gym_name = 'Groove Dance Studio' LIMIT 1));

COMMIT; 