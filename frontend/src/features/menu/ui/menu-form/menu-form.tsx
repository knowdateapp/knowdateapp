import { Flex, IconButton, Menu, MenuButton, MenuItem, MenuList } from '@chakra-ui/react';
import { FC} from 'react';
import { generatePath, useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';
import { Routes } from "shared/config";

export const MenuForm: FC = () => {
  const {t} = useTranslation('translation');
  const navigate = useNavigate();

  // TODO: Как это использовать?
  const onNotesClick = () => navigate(generatePath(Routes.Notes));

  return (
    <Menu>
      <MenuButton
        as={IconButton}
        aria-label="Options"
        variant="outline"
      />
      <MenuList>
        <Flex
          paddingLeft="1em"
          paddingRight="1em"
          flexDirection="column"
          justifyContent="space-between"
          bg="gray.100"
        >
          <MenuItem
            padding="0.5em"
            onClick={onNotesClick}
          >
            {t('features.menu.menu-form.notes')}
          </MenuItem>
          <MenuItem
            padding="0.5em"
          >
            {t('features.menu.menu-form.cards')}
          </MenuItem>
        </Flex>
      </MenuList>
    </Menu>
  );
};
