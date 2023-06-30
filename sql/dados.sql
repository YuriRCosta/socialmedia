INSERT INTO `usuarios` (`id`, `nome`, `nick`, `email`, `senha`, `criadoEm`) VALUES (1, 'Jose', 'Josezito', 'joao@gmail.com', '12345', '2023-06-29 20:44:29');
INSERT INTO `usuarios` (`id`, `nome`, `nick`, `email`, `senha`, `criadoEm`) VALUES (4, 'Fhillipe', 'Felipe', 'felipe@gmail.com', '12345', '2023-06-29 21:00:06');
INSERT INTO `usuarios` (`id`, `nome`, `nick`, `email`, `senha`, `criadoEm`) VALUES (6, 'Fhillipe12', 'Felipe12', 'felip123e@gmail.com', '$2a$10$1yAmeqFDt1ABj0HLlB/xhu7yCPrF8akCAuTlwL/ErO.5973plHBsO', '2023-06-29 22:25:05');
INSERT INTO `usuarios` (`id`, `nome`, `nick`, `email`, `senha`, `criadoEm`) VALUES (7, 'Fhillipe121', 'Felipe121', 'felip1231e@gmail.com', '$2a$10$4TiblohT42wHjzIETRHw8OcogNY7.4ajuuOyVRYrVV5EQOTwl2qmG', '2023-06-29 22:39:46');
INSERT INTO `usuarios` (`id`, `nome`, `nick`, `email`, `senha`, `criadoEm`) VALUES (8, 'Luiza', 'Lu', 'lu@gmail.com', '$2a$10$8NT/ejIjqwm4PmrtwttR1OCUBgHgoEgQDfOs5TaMbm2XBiP0U19jO', '2023-06-30 12:28:03');
INSERT INTO `usuarios` (`id`, `nome`, `nick`, `email`, `senha`, `criadoEm`) VALUES (9, 'Thales', 'Thales', 'thales@gmail.com', '$2a$10$DxyEIuWh63WZTIWER4.ANOpoiUyWZiwjwZzaShjQXmDCOGPlJtdVS', '2023-06-30 12:28:30');

INSERT INTO `seguidores` (`usuario_id`, `seguidor_id`) VALUES (4, 8);
INSERT INTO `seguidores` (`usuario_id`, `seguidor_id`) VALUES (8, 7);
INSERT INTO `seguidores` (`usuario_id`, `seguidor_id`) VALUES (8, 1);
