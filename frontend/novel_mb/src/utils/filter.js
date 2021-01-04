

export default {
  formateTime: (val) => {
    // return val;
    if (val === -62135596800) {
      return '0000-00-00 00:00:00';
    } else {
      const date = new Date(val * 1000);
      const y = date.getFullYear();
      let m = (date.getMonth() + 1).toString();
      m = Number(m) < 10 ? ('0' + m) : m;
      let d = date.getDate().toString();
      d = Number(d) < 10 ? ('0' + d) : d;
      let h = date.getHours().toString();
      h = Number(h) < 10 ? ('0' + h) : h;
      let minute = date.getMinutes().toString();
      let second = date.getSeconds().toString();
      minute = Number(minute) < 10 ? ('0' + minute) : minute;
      second = Number(second) < 10 ? ('0' + second) : second;
      return y + '-' + m + '-' + d + ' ' + h + ':' + minute + ':' + second;
    }
  },
  filterCategory(code) {
    let empty = '';
    let categorySelect = [
      { id: 1, name: '玄幻' },
      { id: 2, name: '修真' },
      { id: 3, name: '都市' },
      { id: 4, name: '穿越' },
      { id: 5, name: '网游' },
      { id: 6, name: '科幻' },
      { id: 7, name: '女频' },
      { id: 0, name: '其他' },
    ]
    categorySelect.forEach((item) => {
      if (item.id === code) {
        empty = item.name;
      }
    });
    return empty;
  },
};
