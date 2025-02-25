import { css } from '@emotion/css';

import { GrafanaTheme2 } from '@grafana/data';
import {
  SceneComponentProps,
  sceneGraph,
  SceneGridRow,
  SceneObjectBase,
  SceneObjectState,
  VizPanel,
} from '@grafana/scenes';
import { Icon, TextLink, useStyles2 } from '@grafana/ui';
import appEvents from 'app/core/app_events';
import { t, Trans } from 'app/core/internationalization';
import { SHARED_DASHBOARD_QUERY } from 'app/plugins/datasource/dashboard/constants';
import { MIXED_DATASOURCE_NAME } from 'app/plugins/datasource/mixed/MixedDataSource';
import { ShowConfirmModalEvent } from 'app/types/events';

import { getDashboardSceneFor, getQueryRunnerFor } from '../../../utils/utils';
import { DashboardScene } from '../../DashboardScene';
import { DashboardGridItem } from '../DashboardGridItem';
import { DefaultGridLayoutManager } from '../DefaultGridLayoutManager';
import { RowRepeaterBehavior } from '../RowRepeaterBehavior';

import { RowOptionsButton } from './RowOptionsButton';

export interface RowActionsState extends SceneObjectState {}

export class RowActions extends SceneObjectBase<RowActionsState> {
  public getParent(): SceneGridRow {
    if (!(this.parent instanceof SceneGridRow)) {
      throw new Error('RowActions must have a SceneGridRow parent');
    }

    return this.parent;
  }

  public getDashboard(): DashboardScene {
    return getDashboardSceneFor(this);
  }

  public removeRow(removePanels?: boolean) {
    const manager = sceneGraph.getAncestor(this, DefaultGridLayoutManager);
    manager.removeRow(this.getParent(), removePanels);
  }

  public onUpdate = (title: string, repeat?: string | null): void => {
    const row = this.getParent();
    let repeatBehavior: RowRepeaterBehavior | undefined;

    if (row.state.$behaviors) {
      for (let b of row.state.$behaviors) {
        if (b instanceof RowRepeaterBehavior) {
          repeatBehavior = b;
        }
      }
    }

    if (title !== row.state.title) {
      row.setState({ title });
    }

    if (repeat) {
      // Remove repeat behavior if it exists
      // to retrigger repeat when adding new one
      if (repeatBehavior) {
        repeatBehavior.removeBehavior();
      }

      repeatBehavior = new RowRepeaterBehavior({ variableName: repeat });
      row.setState({ $behaviors: [...(row.state.$behaviors ?? []), repeatBehavior] });
    } else {
      repeatBehavior?.removeBehavior();
    }
  };

  public onDelete = () => {
    appEvents.publish(
      new ShowConfirmModalEvent({
        title: t('dashboard.default-layout.row-actions.modal.title', 'Delete row'),
        text: t(
          'dashboard.default-layout.row-actions.modal.text',
          'Are you sure you want to remove this row and all its panels?'
        ),
        altActionText: t('dashboard.default-layout.row-actions.modal.alt-action', 'Delete row only'),
        icon: 'trash-alt',
        onConfirm: () => this.removeRow(true),
        onAltAction: () => this.removeRow(),
      })
    );
  };

  public getWarning = () => {
    const row = this.getParent();
    const gridItems = row.state.children;

    const isAnyPanelUsingDashboardDS = gridItems.some((gridItem) => {
      if (!(gridItem instanceof DashboardGridItem)) {
        return false;
      }

      const vizPanel = gridItem.state.body;
      if (vizPanel instanceof VizPanel) {
        const runner = getQueryRunnerFor(vizPanel);
        return (
          runner?.state.datasource?.uid === SHARED_DASHBOARD_QUERY ||
          (runner?.state.datasource?.uid === MIXED_DATASOURCE_NAME &&
            runner?.state.queries.some((query) => query.datasource?.uid === SHARED_DASHBOARD_QUERY))
        );
      }

      return false;
    });

    if (isAnyPanelUsingDashboardDS) {
      return (
        <div>
          <p>
            <Trans i18nKey="dashboard.default-layout.row-actions.repeat.warning.text">
              Panels in this row use the {{ SHARED_DASHBOARD_QUERY }} data source. These panels will reference the panel
              in the original row, not the ones in the repeated rows.
            </Trans>
          </p>
          <TextLink
            external
            href={
              'https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/create-dashboard/#configure-repeating-rows'
            }
          >
            <Trans i18nKey="dashboard.default-layout.row-actions.repeat.warning.learn-more">Learn more</Trans>
          </TextLink>
        </div>
      );
    }

    return undefined;
  };

  static Component = ({ model }: SceneComponentProps<RowActions>) => {
    const dashboard = model.getDashboard();
    const row = model.getParent();
    const { title } = row.useState();
    const { meta, isEditing } = dashboard.useState();
    const styles = useStyles2(getStyles);

    const behaviour = row.state.$behaviors?.find((b) => b instanceof RowRepeaterBehavior);

    return (
      <>
        {meta.canEdit && isEditing && (
          <>
            <div className={styles.rowActions}>
              <RowOptionsButton
                title={title}
                repeat={behaviour instanceof RowRepeaterBehavior ? behaviour.state.variableName : undefined}
                parent={dashboard}
                onUpdate={model.onUpdate}
                warning={model.getWarning()}
              />
              <button
                type="button"
                onClick={model.onDelete}
                aria-label={t('dashboard.default-layout.row-actions.delete', 'Delete row')}
              >
                <Icon name="trash-alt" />
              </button>
            </div>
          </>
        )}
      </>
    );
  };
}

const getStyles = (theme: GrafanaTheme2) => {
  return {
    rowActions: css({
      color: theme.colors.text.secondary,
      lineHeight: '27px',

      button: {
        color: theme.colors.text.secondary,
        paddingLeft: theme.spacing(2),
        background: 'transparent',
        border: 'none',

        '&:hover': {
          color: theme.colors.text.maxContrast,
        },
      },
    }),
  };
};
