import { Workspace } from 'shared/types';
import { apiClient } from 'shared/api';
import { Card, CreateCardParams } from '../types';

const getUrl = (workspace: Workspace) => `${workspace}/cards`;

export const controller = {
  createCard: ({ workspace, question, answer }: CreateCardParams) =>
    apiClient.post<Card>(getUrl(workspace), { question, answer }),
};
