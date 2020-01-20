export function clone(val) {
  return JSON.parse(JSON.stringify(val));
}

export function copy(val) {
  return JSON.parse(JSON.stringify(val));
}
export function dateFormater(date) {
  let year = date.getFullYear().toString(),
    month = (date.getMonth() + 1).toString(),
    day = date.getDate().toString();
  if (month.length === 1) {
    month = `0${month}`;
  }
  if (day.length === 1) {
    day = `0${day}`;
  }
  date = `${year}-${month}-${day}T00:00:00.000Z`;
  return date;
}
// rmvFormatting(val) {
//   let y = new DOMParser().parseFromString(val, "text/html");
//   let x = y.getElementsByTagName("p")[0];
//   return x.innerText;
// },

// clean(content) {
//   return content.replace(/_/g, " ");
// },
// minify(content) {
//   if (content.length == 0) {
//     return content;
//   }
//   content = this.rmvFormatting(content);
//   if (content.length > 70) {
//     let minified = content.substring(0, 70);
//     return minified + "...";
//   }
//   return content;
// },
export function timeFormater(date) {
  let year = date.getFullYear().toString(),
    month = (date.getMonth() + 1).toString(),
    day = date.getDate().toString();
  let hrs = date.getHours().toString(),
    min = date.getMinutes().toString();
  if (year.length === 1) {
    year = `000${year}`;
  }
  if (month.length === 1) {
    month = `0${month}`;
  }
  if (day.length === 1) {
    day = `0${day}`;
  }
  if (hrs.length === 1) {
    hrs = `0${hrs}`;
  }
  if (min.length === 1) {
    min = `0${min}`;
  }
  date = `${year}-${month}-${day}T${hrs}:${min}:00.000`;
  let time = new Date(date);
  return time;
}

export function Notification(
  self,
  message,
  type = "is-success",
  position = "is-top-right",
  duration = 800
) {
  self.$buefy.snackbar.open({
    duration: duration,
    message: message,
    type: type,
    position: position
  });
}
