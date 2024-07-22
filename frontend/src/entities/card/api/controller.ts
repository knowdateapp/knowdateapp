import { Workspace } from 'shared/model';
import { apiClient } from 'shared/api';
import { CardEntity, CreateCardParams } from '../types';

const getUrl = (workspace: Workspace) => `${workspace}/cards`;

export const controller = {
  createCard: ({ workspace, question, answer }: CreateCardParams) =>
    apiClient.post<CardEntity>(getUrl(workspace), { question, answer }),
};
