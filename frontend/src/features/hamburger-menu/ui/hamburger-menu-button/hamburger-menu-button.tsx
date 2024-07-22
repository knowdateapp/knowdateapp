import { Flex, IconButton, Menu, MenuItem, MenuList, MenuButton } from '@chakra-ui/react';
import { HamburgerIcon } from '@chakra-ui/icons';
import { FC } from 'react';
import { generatePath, useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';
import { Routes } from 'shared/config';

export const HamburgerMenuButton: FC = () => {
  const { t } = useTranslation('translation');
  const navigate = useNavigate();

  const onNotesClick = () => navigate(generatePath(Routes.Notes));
  const onCardsClick = () => navigate(generatePath(Routes.Cards));

  return (
    // TODO: Проверить стили.
    <Menu>
      <MenuButton as={IconButton} icon={<HamburgerIcon />} aria-label="Options" variant="outline" />
      <MenuList>
        <Flex
          flexDirection="column"
          justifyContent="space-between"
          gap={2}
          minW={200}
          padding="0.5em 0"
          bg="gray.100"
        >
          <MenuItem justifyContent="center" onClick={onNotesClick}>
            {t('features.menu.menuForm.notes')}
          </MenuItem>
          <MenuItem justifyContent="center" onClick={onCardsClick}>
            {t('features.menu.menuForm.cards')}
          </MenuItem>
        </Flex>
      </MenuList>
    </Menu>
  );
};
