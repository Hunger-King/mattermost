// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import {FormattedMessage, useIntl} from 'react-intl';

import type {UserReport} from '@mattermost/types/reports';

import type {Column, CoreColumn} from 'components/admin_console/list_table';
import * as Menu from 'components/menu';
import Input from 'components/widgets/inputs/input/input';

import {ColumnNames} from '../constants';

import './system_users_column_toggler.scss';

interface Props {
    allColumns: Array<Column<UserReport>>;
    visibleColumnsLength: number;
}

export function SystemUsersColumnTogglerMenu(props: Props) {
    const {formatMessage} = useIntl();

    function getColumnName(columnId: CoreColumn<UserReport, unknown>['id']) {
        switch (columnId) {
        case ColumnNames.username:
            return (
                <FormattedMessage
                    id='admin.system_users.column_toggler.column.displayName'
                    defaultMessage='Display Name'
                />
            );
        case ColumnNames.email:
            return (
                <FormattedMessage
                    id='admin.system_users.column_toggler.column.email'
                    defaultMessage='Email'
                />
            );
        case ColumnNames.createAt:
            return (
                <FormattedMessage
                    id='admin.system_users.column_toggler.column.createAt'
                    defaultMessage='Create At'
                />
            );
        case ColumnNames.lastLoginAt:
            return (
                <FormattedMessage
                    id='admin.system_users.column_toggler.column.lastLoginAt'
                    defaultMessage='Last Login At'
                />
            );
        case ColumnNames.lastStatusAt:
            return (
                <FormattedMessage
                    id='admin.system_users.column_toggler.column.lastStatusAt'
                    defaultMessage='Last Status At'
                />
            );
        case ColumnNames.lastPostDate:
            return (
                <FormattedMessage
                    id='admin.system_users.column_toggler.column.lastPostDate'
                    defaultMessage='Last Post Date'
                />
            );
        case ColumnNames.daysActive:
            return (
                <FormattedMessage
                    id='admin.system_users.column_toggler.column.daysActive'
                    defaultMessage='Days Active'
                />
            );
        case ColumnNames.totalPosts:
            return (
                <FormattedMessage
                    id='admin.system_users.column_toggler.column.totalPosts'
                    defaultMessage='Total Posts'
                />
            );
        case ColumnNames.actions:
            return (
                <FormattedMessage
                    id='admin.system_users.column_toggler.column.actions'
                    defaultMessage='Actions'
                />
            );
        default:
            return <span/>;
        }
    }

    return (
        <div className='systemUsersColumnToggler'>
            <Menu.Container
                menuButton={{
                    id: 'systemUsersColumnTogglerMenuButton',
                    class: 'inputWithMenu',
                    'aria-label': formatMessage({
                        id: 'admin.system_users.column_toggler.menuButtonAriaLabel',
                        defaultMessage:
                            'Open menu to select columns to display',
                    }),
                    as: 'div',
                    children: (
                        <Input
                            label={formatMessage({
                                id: 'admin.system_users.column_toggler.placeholder',
                                defaultMessage: 'Columns',
                            })}
                            name='colXC'
                            value={formatMessage(
                                {
                                    id: 'admin.system_users.column_toggler.menuButtonText',
                                    defaultMessage: '{selectedCount} selected',
                                },
                                {
                                    selectedCount: props.visibleColumnsLength,
                                },
                            )}
                            readOnly={true}
                            inputSuffix={
                                <i className='icon icon-chevron-down'/>
                            }
                        />
                    ),
                }}
                menu={{
                    id: 'systemUsersColumnTogglerMenu',
                    'aria-label': formatMessage({
                        id: 'admin.system_users.column_toggler.dropdownAriaLabel',
                        defaultMessage: 'Columns visibility menu',
                    }),
                }}
            >
                {props.allColumns.map((column) => {
                    let leadingElement;
                    if (column.getCanHide()) {
                        if (column.getIsVisible()) {
                            leadingElement = (
                                <i className='icon icon-checkbox-marked'/>
                            );
                        } else {
                            leadingElement = (
                                <i className='icon icon-checkbox-blank-outline'/>
                            );
                        }
                    } else {
                        // This means the column is always visible
                        leadingElement = (
                            <i className='icon icon-checkbox-marked'/>
                        );
                    }

                    return (
                        <Menu.Item
                            key={column.id}
                            id={column.id}
                            role='menuitemcheckbox'
                            labels={getColumnName(column.id)}
                            disabled={!column.getCanHide()}
                            leadingElement={leadingElement}
                            onClick={column.getToggleVisibilityHandler()}
                        />
                    );
                })}
            </Menu.Container>
        </div>
    );
}