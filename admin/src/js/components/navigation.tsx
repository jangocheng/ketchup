import debounce from 'lodash-es/debounce';
import msx from 'lib/msx';
import * as m from 'mithril';
import { AuthController } from 'components/auth';
import Button from 'components/button';
let store = require('store/dist/store.modern') as StoreJSStatic;


export default class NavigationComponent extends AuthController {
  collapsed: boolean;
  constructor() {
    super();
    this.collapsed = store.get('hideMenu') || window.innerWidth <= 480;
    window.addEventListener('resize', this.resizeHandler);
    this.resizeHandler();
  }

  onunload() {
    window.removeEventListener('resize', this.resizeHandler);
  }

  resizeHandler = debounce(() => {
    if (window.innerWidth > 480) {
      return;
    }
    store.set('hideMenu', true);
    this.collapsed = true;
    m.redraw();
  }, 300);

  toggle() {
    let collapsed = !this.collapsed;
    store.set('hideMenu', collapsed);
    this.collapsed = collapsed;
  }

  link(url: string, text: string, opts: { onclick?: () => void, additionalClasses?: string, icon?: string } = {}) {
    return <a class={`nav-link ${opts.additionalClasses}`}
      href={url}
      oncreate={m.route.link}
      onclick={opts.onclick}
    >
      {!!opts.icon ? <span class={`typcn typcn-${opts.icon}`}></span> : ''}
      <span class='nav-link__text'>{text}</span>
    </a>;
  }

  view() {
    let navClass = 'navigation';
    if (this.collapsed) {
      navClass += ' navigation--hidden';
    }

    if (!this.user) {
      return <div class={navClass}>
        {this.link('/admin', 'K', { additionalClasses: 'nav-title' })}
        {this.link('/admin/login', 'Login')}
      </div>;
    }

    return <div class={navClass}>
      {this.link('/admin', 'K', { additionalClasses: 'nav-title' })}
      <div class='nav-button'>
        <Button
          class='button--green button--center'
          href='/admin/compose'
        >
          <span class='typcn typcn-edit' />
          <span class='nav-link__text'>Compose</span>
        </Button>
      </div>
      {this.link('/admin/pages', 'Pages', { icon: 'document-text' })}
      {this.link('/admin/themes', 'Theme', { icon: 'brush' })}
      {this.link('/admin/data', 'Data', { icon: 'th-small' })}
      {this.link('/admin/settings', 'Settings', { icon: 'spanner-outline' })}
      {this.link('/admin/logout', 'Logout', { onclick: () => this.logout(), icon: 'weather-night' })}
      <a class='nav-link nav-link--toggle' onclick={() => this.toggle()}>
        <span class={`typcn typcn-arrow-${this.collapsed ? 'maximise' : 'minimise'}`} />
      </a>
    </div>;
  }
}
