CREATE TABLE IF NOT EXISTS knowledgebase (id SERIAL PRIMARY KEY,topic VARCHAR(1000) UNIQUE NOT NULL, answers VARCHAR(1000) NOT NULL);
CREATE TABLE IF NOT EXISTS samplequestion (id SERIAL PRIMARY KEY,question VARCHAR(1000) UNIQUE NOT NULL);

-----------Insert KnowledgeBase-----------------
INSERT INTO knowledgebase (id, topic , answers) VALUES (1, 'What are your products made of', 'Our aim is to make products using only renewable resources or recycled materials — so we prioritize, responsibly source, and recover materials.');

INSERT INTO knowledgebase (id, topic , answers) VALUES (2, 'Where is my order', 'You can check order status, track a delivery, view pickup details, edit your delivery or email address, print an invoice, and more by checking your online Order Status');

INSERT INTO knowledgebase (id, topic , answers) VALUES (3, 'I need to make a return. What is your refund policy', 'Only items that have been purchased directly  either online or at an Retail Store, can be returned. Our products purchased through other retailers must be returned in accordance with their respective returns and refunds policy.');

INSERT INTO knowledgebase (id, topic , answers) VALUES (4, 'Where is your store', 'Use the Best Buy store locator to find stores in your area.');

INSERT INTO knowledgebase (id, topic , answers) VALUES (5, 'do you offer any discounts or coupons ', 'Discounts are sometimes offered when you buy more than one of a particular item. Some offers and coupons are only valid in stores.');

INSERT INTO knowledgebase (id, topic , answers) VALUES (6, 'I Have a Problem What Do I Do', 'If you need help, we offer a number of options for customer support.');

----------Insert SampleQuestions------------------

INSERT INTO samplequestion (id, question ) VALUES (1, 'What are your products made of?');

INSERT INTO samplequestion (id, question ) VALUES (2, 'Where is my order?');

INSERT INTO samplequestion (id, question ) VALUES (3, 'What is your refund policy?');



